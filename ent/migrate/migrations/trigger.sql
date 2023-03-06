-- Trigger to keep financial_accounts' cards in-sync.
--
-- NOTE: One-way sync from cards nullable parent table (financial_account) to avoid infinite cyclic-recursive calls.
-- Both backend and frontend system MUST ensure none of these update fields are modified IF card has a linked fin_account.
CREATE OR REPLACE FUNCTION sync_fin_account_debit_card() RETURNS TRIGGER
AS $$
BEGIN
UPDATE cards
SET balance = new.balance,global_balance = new.global_balance,
    bank_name = new.bank_name,currency_code = new.currency_code,
    update_time = (NOW() AT TIME ZONE 'utc'),row_version = row_version+1,last_update_by = 'system'
WHERE linked_financial_account = new.account_id;
RETURN NEW;
END;
$$
LANGUAGE PLPGSQL;

CREATE OR REPLACE TRIGGER sync_fin_account_debit_card_tgr AFTER UPDATE ON financial_accounts FOR EACH ROW EXECUTE PROCEDURE sync_fin_account_debit_card();

-- Trigger to populate common fields from cards parent table (financial_account) with default values.
-- Furthermore, if linked financial account is not null, this trigger will override incoming data with parent values to enhance table in-sync mechanism.
CREATE OR REPLACE FUNCTION set_cards_parent_values() RETURNS TRIGGER
AS $$
BEGIN
	IF NEW.linked_financial_account IS NOT NULL THEN
SELECT INTO
    NEW.balance,NEW.global_balance,NEW.currency_code,NEW.bank_name
    balance,global_balance,currency_code,bank_name
FROM financial_accounts WHERE account_id = NEW.linked_financial_account;
END IF;
RETURN NEW;
END;
$$
LANGUAGE PLPGSQL;

CREATE OR REPLACE TRIGGER set_cards_parent_values_tgr BEFORE INSERT ON cards FOR EACH ROW EXECUTE PROCEDURE set_cards_parent_values();


-- Trigger to populate user_financial_stats table initial values ON user creation.
CREATE OR REPLACE FUNCTION create_user_financial_stats() RETURNS TRIGGER
AS $$
BEGIN
INSERT INTO user_financial_stats(user_id,currency_code) VALUES (NEW.user_id,NEW.global_currency_code);
RETURN NEW;
END;
$$
LANGUAGE PLPGSQL;

CREATE OR REPLACE TRIGGER create_user_financial_stats_tgr AFTER INSERT ON users FOR EACH ROW EXECUTE PROCEDURE create_user_financial_stats();

-- FIXME: To be replaced using stream processing
-- Trigger to populate user_financial_stats table.
-- Feeding the given table AT WRITE time WILL improve performance and avoid runtime aggregation functions calls AT READ TIME.
-- CREATE OR REPLACE FUNCTION update_user_financial_stats_from_card() RETURNS TRIGGER
-- AS $$
-- DECLARE
-- total_balance_sum MONEY;
-- 	total_debt_sum MONEY;
-- BEGIN
-- SELECT SUM(global_balance) INTO total_debt_sum AS currency_code FROM cards WHERE card_type = 'CREDIT' AND owner_id = NEW.owner_id AND cards.is_active = true;
-- SELECT SUM(global_balance - total_debt_sum) INTO total_balance_sum FROM cards WHERE card_type = 'DEBIT' AND owner_id = NEW.owner_id AND cards.is_active = true AND linked_financial_account IS NOT NULL;
-- UPDATE user_financial_stats SET
--                                 total_balance = total_balance_sum,total_debt = total_debt_sum,
--                                 update_time = (NOW() AT TIME ZONE 'utc'),row_version = row_version+1,last_update_by = 'system'
-- WHERE user_id = NEW.owner_id;
-- RETURN NEW;
-- END;
-- $$
-- LANGUAGE PLPGSQL;
--
-- CREATE OR REPLACE TRIGGER refresh_user_financial_stats_from_card_tgr AFTER INSERT OR UPDATE ON cards FOR EACH ROW EXECUTE PROCEDURE update_user_financial_stats_from_card();
