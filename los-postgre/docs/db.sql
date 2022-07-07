CREATE TABLE users (
	id VARCHAR(200) PRIMARY KEY,
	username VARCHAR(200) NOT NULL UNIQUE,
	password VARCHAR(200) NOT NULL,
	is_officer BOOLEAN DEFAULT false,
	created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE loan_applications (
	id VARCHAR(200) PRIMARY KEY,
	user_id VARCHAR(200) NOT NULL REFERENCES users(id),
	officer_id VARCHAR(200) REFERENCES users(id),
	full_name VARCHAR(200) DEFAULT '',
	birth_date VARCHAR(200) DEFAULT '',
	full_address VARCHAR(200) DEFAULT '',
	phone VARCHAR(200) DEFAULT '',
	id_card_url VARCHAR(200) DEFAULT '',
	other_business VARCHAR(200) DEFAULT '',
	status VARCHAR(25) DEFAULT '',
	is_private_field BOOLEAN DEFAULT false,
	exp_in_year SMALLINT DEFAULT 0,
	active_field_number SMALLINT DEFAULT 0,
	sow_seeds_per_cycle SMALLINT DEFAULT 0,
	needed_fertilizier_per_cycle_in_kg SMALLINT DEFAULT 0,
	estimated_yield_in_kg SMALLINT DEFAULT 0,
	estimated_price_of_harvest_per_kg SMALLINT DEFAULT 0,
	harvest_cycle_in_months SMALLINT DEFAULT 0,
	loan_application_in_idr BIGINT DEFAULT 0,
	business_income_per_month_in_idr BIGINT DEFAULT 0,
	business_outcome_per_month_in_idr BIGINT DEFAULT 0,
	created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);