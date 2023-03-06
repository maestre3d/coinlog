-- NOTE: A global currency code is used to perform currency exchange calculations from user's financial telemetry pipeline.
-- Furthermore, currencies are IMMUTABLES to avoid any further due while performing statistic/aggregation calculations.
CREATE TABLE IF NOT EXISTS users(
	user_id BIGSERIAL PRIMARY KEY,
	display_name VARCHAR(64) NOT NULL,
	username VARCHAR(48) NOT NULL UNIQUE,
	image_url VARCHAR(2048),
	locale_code VARCHAR(10) DEFAULT 'en_US',
	global_currency_code VARCHAR(6) NOT NULL,
	create_time TIMESTAMP WITHOUT TIME ZONE DEFAULT (NOW() AT TIME ZONE 'utc'),
	update_time TIMESTAMP WITHOUT TIME ZONE DEFAULT (NOW() AT TIME ZONE 'utc'),
	is_active BOOLEAN DEFAULT TRUE,
	last_update_by VARCHAR(128) DEFAULT 'system',
	row_version BIGINT DEFAULT 1
);

CREATE TABLE IF NOT EXISTS contacts(
	contact_id BIGSERIAL PRIMARY KEY,
	display_name VARCHAR(64) NOT NULL,
	image_url VARCHAR(2048),
	create_time TIMESTAMP WITHOUT TIME ZONE DEFAULT (NOW() AT TIME ZONE 'utc'),
	update_time TIMESTAMP WITHOUT TIME ZONE DEFAULT (NOW() AT TIME ZONE 'utc'),
	is_active BOOLEAN DEFAULT TRUE,
	last_update_by VARCHAR(128) DEFAULT 'system',
	row_version BIGINT DEFAULT 1,
	owner_id BIGINT NOT NULL,
	linked_to BIGINT,
	CONSTRAINT user_contact FOREIGN KEY(owner_id) REFERENCES users(user_id) ON UPDATE CASCADE ON DELETE CASCADE,
	CONSTRAINT user_linked FOREIGN KEY(linked_to) REFERENCES users(user_id) ON UPDATE CASCADE ON DELETE SET NULL
);

-- NOTE: Global MONEY fields (e.g. global_balance) are used by the system to ensure a specific exchange conversion for
-- data processing pipeline calculations. Otherwise, performing statistic functions would be compute-expensive requiring the data pipeline process to
-- do the exchange conversion for each item AT READ TIME.
-- Moreover, these fields should be computed and then populated by the backend system (might require an external call to currency_exchange service) AT WRITE TIME.
-- The backend system MUST use users.global_currency_code field to use it as currency exchange type.
--
-- NOTE: Backend system COULD set default currency_code value using users.global_currency_code data AT WRITE TIME.
CREATE TABLE IF NOT EXISTS financial_accounts(
	account_id BIGSERIAL PRIMARY KEY,
	display_name VARCHAR(64) NOT NULL,
	account_type VARCHAR(64) NOT NULL,
	last_digits SMALLINT,
	bank_name VARCHAR(48),
	balance MONEY DEFAULT 0.0,
	global_balance MONEY DEFAULT 0.0,
	currency_code VARCHAR(6) NOT NULL,
	create_time TIMESTAMP WITHOUT TIME ZONE DEFAULT (NOW() AT TIME ZONE 'utc'),
	update_time TIMESTAMP WITHOUT TIME ZONE DEFAULT (NOW() AT TIME ZONE 'utc'),
	is_active BOOLEAN DEFAULT TRUE,
	last_update_by VARCHAR(128) DEFAULT 'system',
	row_version BIGINT DEFAULT 1,
	owner_id BIGINT NOT NULL,
	CONSTRAINT user_financial_account FOREIGN KEY(owner_id) REFERENCES users(user_id) ON UPDATE CASCADE ON DELETE CASCADE
);

-- NOTE: Global MONEY fields (e.g. global_balance, global_loan_limit) are used by the system to ensure a specific exchange conversion for
-- data processing pipeline calculations. Otherwise, performing statistic functions would be compute-expensive requiring the data pipeline process to
-- do the exchange conversion for each item at read time.
-- Moreover, these fields should be computed and then populated by the backend system (might require an external call to currency_exchange service) AT WRITE TIME.
-- The backend system MUST use users.global_currency_code field to use it as currency exchange type.
--
-- NOTE: Backend system COULD set default currency_code value using users.global_currency_code data AT WRITE TIME.
--
-- NOTE: A CREDIT card type MUST NOT have a financial_account link.
--
-- NOTE: Backend system MUST check if there is a currency_code value mismatch between this table and linked fin_account (if linked fin_account is not null).
CREATE TABLE IF NOT EXISTS cards(
	card_id BIGSERIAL PRIMARY KEY,
	display_name VARCHAR(64),
	card_type VARCHAR(48),
	bank_name VARCHAR(48),
	last_digits SMALLINT,
	cutoff_last_date DATE,
	cutoff_interval_days SMALLINT DEFAULT 0,
	cutoff_interval_months SMALLINT DEFAULT 0,
	cutoff_interval_years SMALLINT DEFAULT 0,
	balance MONEY DEFAULT 0.0,
	loan_limit MONEY DEFAULT 0.0,
	global_balance MONEY DEFAULT 0.0,
	global_loan_limit MONEY DEFAULT 0.0,
	currency_code VARCHAR(6) NOT NULL,
	create_time TIMESTAMP WITHOUT TIME ZONE DEFAULT (NOW() AT TIME ZONE 'utc'),
	update_time TIMESTAMP WITHOUT TIME ZONE DEFAULT (NOW() AT TIME ZONE 'utc'),
	is_active BOOLEAN DEFAULT TRUE,
	last_update_by VARCHAR(128) DEFAULT 'system',
	row_version BIGINT DEFAULT 1,
	owner_id BIGINT NOT NULL,
	linked_financial_account BIGINT,
	CONSTRAINT user_card FOREIGN KEY(owner_id) REFERENCES users(user_id) ON UPDATE CASCADE ON DELETE CASCADE,
	CONSTRAINT financial_account_card FOREIGN KEY(linked_financial_account) REFERENCES financial_accounts(account_id) ON UPDATE CASCADE ON DELETE SET NULL
);

-- Trigger to keep financial_accounts' cards in-sync.
--
-- NOTE: One-way sync from cards nullable parent table (financial_account) to avoid infinite recursive calls.
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

CREATE TABLE IF NOT EXISTS user_financial_stats(
	user_id BIGINT PRIMARY KEY,
	total_balance MONEY DEFAULT 0.0,
	total_debt MONEY DEFAULT 0.0,
	currency_code VARCHAR(6) NOT NULL,
	create_time TIMESTAMP WITHOUT TIME ZONE DEFAULT (NOW() AT TIME ZONE 'utc'),
	update_time TIMESTAMP WITHOUT TIME ZONE DEFAULT (NOW() AT TIME ZONE 'utc'),
	is_active BOOLEAN DEFAULT TRUE,
	last_update_by VARCHAR(128) DEFAULT 'system',
	row_version BIGINT DEFAULT 1,
	CONSTRAINT user_fin_stats FOREIGN KEY(user_id) REFERENCES users(user_id) ON UPDATE CASCADE ON DELETE CASCADE
);

-- Trigger to populate user_financial_stats table initial values ON user creation.
CREATE OR REPLACE FUNCTION create_user_financial_stats() RETURNS TRIGGER
AS $$
DECLARE
	total_balance_sum MONEY;
	total_debt_sum MONEY;
BEGIN
	INSERT INTO user_financial_stats(user_id,currency_code) VALUES (NEW.user_id,NEW.global_currency_code);
	RETURN NEW;
END;
$$
LANGUAGE PLPGSQL;

CREATE OR REPLACE TRIGGER create_user_financial_stats_tgr AFTER INSERT ON users FOR EACH ROW EXECUTE PROCEDURE create_user_financial_stats();

-- Trigger to populate user_financial_stats table.
-- Feeding the given table AT WRITE time WILL improve performance and avoid runtime aggregation functions calls AT READ TIME.
CREATE OR REPLACE FUNCTION update_user_financial_stats_from_card() RETURNS TRIGGER
AS $$
DECLARE
	total_balance_sum MONEY;
	total_debt_sum MONEY;
BEGIN
	SELECT SUM(global_balance) INTO total_debt_sum AS currency_code FROM cards WHERE card_type = 'CREDIT' AND owner_id = NEW.owner_id AND cards.is_active = true;
	SELECT SUM(global_balance - total_debt_sum) INTO total_balance_sum FROM cards WHERE card_type = 'DEBIT' AND owner_id = NEW.owner_id AND cards.is_active = true AND linked_financial_account IS NOT NULL;
	UPDATE user_financial_stats SET
		total_balance = total_balance_sum,total_debt = total_debt_sum,
		update_time = (NOW() AT TIME ZONE 'utc'),row_version = row_version+1,last_update_by = 'system'
		WHERE user_id = NEW.owner_id;
	RETURN NEW;
END;
$$
LANGUAGE PLPGSQL;

CREATE OR REPLACE TRIGGER refresh_user_financial_stats_from_card_tgr AFTER INSERT OR UPDATE ON cards FOR EACH ROW EXECUTE PROCEDURE update_user_financial_stats_from_card();

select * from cards JOIN users ON user_id = owner_id;

-- SELECT SUM(global_balance) AS currency_code FROM cards WHERE card_type = 'CREDIT' AND owner_id = 1 AND cards.is_active = true;

SELECT SUM(global_balance - debt.total_debt) AS total_balance_sum FROM cards, (SELECT SUM(global_balance) AS total_debt FROM cards WHERE card_type = 'CREDIT' AND owner_id = 1 AND cards.is_active = true) AS debt
WHERE card_type = 'DEBIT' AND owner_id = 1 AND cards.is_active = true;

CREATE TABLE IF NOT EXISTS financial_transactions(
	transaction_id BIGSERIAL PRIMARY KEY,
	display_name VARCHAR(96),
	amount MONEY NOT NULL,
	global_amount MONEY NOT NULL,
	currency_code VARCHAR(6) NOT NULL,
	create_time TIMESTAMP WITHOUT TIME ZONE DEFAULT (NOW() AT TIME ZONE 'utc'),
	update_time TIMESTAMP WITHOUT TIME ZONE DEFAULT (NOW() AT TIME ZONE 'utc'),
	is_active BOOLEAN DEFAULT TRUE,
	last_update_by VARCHAR(128) DEFAULT 'system',
	row_version BIGINT DEFAULT 1,
	owner_id BIGINT NOT NULL,
	contact_id BIGINT,
	card_id BIGINT,
	finance_account_id BIGINT,
	CONSTRAINT user_fin_transactions FOREIGN KEY owner_id REFERENCES users(user_id) ON UPDATE CASCADE ON DELETE CASCADE,
	CONSTRAINT contact_fin_transactions FOREIGN KEY contact_id REFERENCES contacts(contact_id) ON UPDATE CASCADE ON DELETE SET NULL,
	CONSTRAINT card_fin_transactions FOREIGN KEY card_id REFERENCES cards(card_id) ON UPDATE CASCADE ON DELETE SET NULL,
	CONSTRAINT fin_acc_fin_transactions FOREIGN KEY finance_account_id REFERENCES finance_accounts ON UPDATE CASCADE ON DELETE SET NULL
);

SELECT * FROM user_financial_stats AS fs RIGHT JOIN users ON users.user_id = fs.user_id;

-- TODO: Finish finance_tx table schema definition
-- TODO: Add trigger to update card or finance account balance when insert/update a fin_tx.
