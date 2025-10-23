CREATE TABLE customers (
   id SERIAL PRIMARY KEY,
   name VARCHAR(100) NOT NULL,
   email VARCHAR(100) UNIQUE NOT NULL,
   type VARCHAR(20) DEFAULT 'INDIVIDUAL',
   status VARCHAR(20) DEFAULT 'ACTIVE',
   company_name VARCHAR(255),
   premium_tier VARCHAR(50),
   
   -- Personal info for individual customers
   phone VARCHAR(20),
   address TEXT,
   date_of_birth DATE,
   
   -- Business info for business customers
   tax_id VARCHAR(50),
   industry VARCHAR(100),
   employee_count INT,
   website VARCHAR(255),
   
   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes for better query performance
CREATE INDEX idx_customers_type ON customers(type);
CREATE INDEX idx_customers_status ON customers(status);
CREATE INDEX idx_customers_email ON customers(email);
CREATE INDEX idx_customers_premium_tier ON customers(premium_tier);