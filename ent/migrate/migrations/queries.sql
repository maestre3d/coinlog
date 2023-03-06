-- FETCH ALL CONTACTS OF A SPECIFIC USER WITH LINKED_USER POPULATED IF NOT NULL
SELECT * FROM contacts LEFT JOIN users ON contacts.linked_to = users.user_id WHERE contacts.owner_id = 1;

-- FETCH ALL CARDS OF A SPECIFIC USER
SELECT * FROM cards WHERE owner_id = 2 AND is_active = true;

-- GET TOTAL BALANCE OF A SPECIFIC USER
SELECT SUM(global_balance - loans.debt) AS total_balance, global_currency_code AS currency_code FROM cards, (SELECT SUM(global_balance) AS debt FROM cards WHERE card_type = 'CREDIT' AND owner_id = 2 AND is_active = true) AS loans, users
WHERE card_type = 'DEBIT' AND owner_id = 2 AND cards.is_active = true AND user_id = 2 GROUP BY global_currency_code;

-- GET TOTAL DEBT OF A SPECIFIC USER
SELECT SUM(global_balance) AS total_debt, global_currency_code AS currency_code FROM cards, users WHERE card_type = 'CREDIT' AND owner_id = 2 AND cards.is_active = true AND user_id = 2 GROUP BY global_currency_code;
