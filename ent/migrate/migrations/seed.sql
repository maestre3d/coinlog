INSERT INTO users(display_name, username, image_url, global_currency_code) VALUES
('Alonso Ruiz','aruiz','https://static.coinlog.info/users/aruiz.jpg','MXN'),
('Bruno Arevalo','xerjes.dicaprio','https://static.coinlog.info/users/xerjes-dicaprio.jpeg','MXN'),
('Victor Gramer','vicmolgram','https://static.coinlog.info/users/vicmolgram.png','MXN'),
('Fernando Herrera','fernandoha',NULL,'USD');

INSERT INTO contacts(display_name, owner_id, linked_to) VALUES
('Brunacido',1,2),('Vicmolgram',1,3),
('Alfonso Zaragoza',2,NULL),('Fernando H',1,NULL),('Jose Rivera',3,NULL);

INSERT INTO financial_accounts(display_name,account_type,last_digits,bank_name,balance,global_balance,currency_code,owner_id) VALUES
('Payroll','CHECKING',876,'Citibanamex',156530,156530,'MXN',1),('RappiCard','CHECKING',723,'Rappi',524.76,524.76,'MXN',1);

-- DEBIT
INSERT INTO cards(display_name,card_type,owner_id,balance,global_balance,currency_code,linked_financial_account) VALUES
('Payroll','DEBIT',1,156530,156530,'MXN',1),('Rappi Account','DEBIT',1,524.76,524.76,'MXN',2),
('Scotiabank Daddy','DEBIT',2,7834.40,145194.15,'USD',NULL),
('Banorte Payroll','DEBIT',3,3560,3650,'MXN',NULL);
-- CREDIT
INSERT INTO cards(display_name, card_type, owner_id, balance, loan_limit, global_balance, global_loan_limit, currency_code, cutoff_last_date, cutoff_interval_months) VALUES
('BBVA Blue','CREDIT',1,10300,60000,10300,60000,'MXN',DATE('2022-12-31'),1),('Rappi Black','CREDIT',1,5400,12000,5400,12000,'MXN',DATE('2023-01-14'),1),
('American Express Gold','CREDIT',2,15040.50,100000,15040.50,100000,'MXN',DATE('2023-02-08'),1),
('BBVA goldie','CREDIT',4,39760,120000,2145.41,6475.08,'MXN',DATE('2022-12-05'),1);