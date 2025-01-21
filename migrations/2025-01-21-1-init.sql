CREATE TABLE vacancies (
	id SERIAL PRIMARY KEY,
	email VARCHAR(255) NOT NULL,
	role VARCHAR(255),
	company VARCHAR(255),
	salary VARCHAR(255),
	type VARCHAR(255),
	location VARCHAR(255)
)