# Data Model: Rental Property Management Platform

## Entity Relationship Overview

```
User ||--o{ Property : owns
Property ||--o{ Comment : has
Property ||--|| FinancialMetrics : calculates
Property ||--o{ PropertyValuation : includes
User ||--o{ BuyingBoxCriteria : defines
```

## Core Entities

### User
**Purpose**: Represents individual investors and team members who use the platform

**Fields**:
- `id` (UUID, Primary Key): Unique identifier
- `email` (String, Unique, NOT NULL): User login email
- `password_hash` (String, NOT NULL): Bcrypt hashed password
- `first_name` (String, NOT NULL): User's first name
- `last_name` (String, NOT NULL): User's last name
- `created_at` (Timestamp): Account creation time
- `updated_at` (Timestamp): Last profile update
- `is_active` (Boolean, Default: true): Account status

**Validation Rules**:
- Email must be valid format and unique
- Password minimum 8 characters
- Names required, max 50 characters each

**State Transitions**:
- Created → Active (default)
- Active → Inactive (admin action)
- Inactive → Active (reactivation)

### Property
**Purpose**: Represents rental properties with all investment-related data

**Fields**:
- `id` (UUID, Primary Key): Unique identifier
- `user_id` (UUID, Foreign Key): Owner reference
- `address` (String, NOT NULL): Full property address
- `year_built` (Integer): Construction year
- `land_area_sqft` (Integer): Land area in square feet
- `building_area_sqft` (Integer): Building area in square feet
- `purchase_price` (Decimal(12,2)): Property purchase price
- `intended_rent` (Decimal(10,2)): Target monthly rent
- `created_at` (Timestamp): Record creation time
- `updated_at` (Timestamp): Last modification time

**JSON Fields** (PostgreSQL JSONB):
- `operating_expenses` (JSON): Insurance, HOA, taxes, utilities
- `financing_terms` (JSON): Interest rate, loan term, down payment, closing costs
- `operating_assumptions` (JSON): Vacancy rate, maintenance %, management fees
- `local_context` (JSON): School scores, livability scores

**Validation Rules**:
- Address required, max 255 characters
- Purchase price must be positive
- Year built between 1800 and current year + 1
- Areas must be positive integers

**Indexes**:
- `idx_property_user_id` on user_id
- `idx_property_address` on address (GIN for search)
- `idx_property_purchase_price` on purchase_price

### PropertyValuation
**Purpose**: Third-party valuation data from Zillow, Redfin, etc.

**Fields**:
- `id` (UUID, Primary Key): Unique identifier
- `property_id` (UUID, Foreign Key): Property reference
- `source` (String, NOT NULL): Valuation source (Zillow, Redfin, Rentimate)
- `valuation_type` (String, NOT NULL): 'market_value' or 'rental_estimate'
- `value` (Decimal(12,2), NOT NULL): Estimated value
- `valuation_date` (Date): Date of valuation
- `created_at` (Timestamp): Record creation time

**Validation Rules**:
- Source must be one of: Zillow, Redfin, Rentimate
- Valuation type must be: market_value, rental_estimate
- Value must be positive
- Valuation date cannot be in future

**Indexes**:
- `idx_valuation_property_id` on property_id
- `idx_valuation_source_type` on (source, valuation_type)

### FinancialMetrics
**Purpose**: Calculated investment metrics for each property

**Fields**:
- `id` (UUID, Primary Key): Unique identifier
- `property_id` (UUID, Foreign Key): Property reference
- `monthly_mortgage_payment` (Decimal(10,2)): Calculated mortgage payment
- `net_operating_income` (Decimal(10,2)): Annual NOI
- `cap_rate` (Decimal(5,2)): Cap rate percentage
- `cash_on_cash_return` (Decimal(5,2)): CoC return percentage
- `cash_to_close` (Decimal(12,2)): Total cash needed
- `rent_to_value_ratio` (Decimal(5,2)): RTV percentage
- `gross_rent_multiplier` (Decimal(5,2)): GRM value
- `calculated_at` (Timestamp): Calculation timestamp
- `is_current` (Boolean): Whether calculation is up-to-date

**Validation Rules**:
- All decimal values must be finite
- Percentages between 0 and 100
- Cash amounts must be positive

**Indexes**:
- `idx_metrics_property_id` on property_id
- `idx_metrics_cap_rate` on cap_rate
- `idx_metrics_coc_return` on cash_on_cash_return

### Comment
**Purpose**: User collaboration and notes on properties

**Fields**:
- `id` (UUID, Primary Key): Unique identifier
- `property_id` (UUID, Foreign Key): Property reference
- `user_id` (UUID, Foreign Key): Comment author
- `content` (Text, NOT NULL): Comment text
- `created_at` (Timestamp): Comment creation time
- `updated_at` (Timestamp): Last edit time
- `parent_id` (UUID, Foreign Key): For threaded comments (optional)

**Validation Rules**:
- Content required, max 2000 characters
- Content must not be empty after trimming

**Indexes**:
- `idx_comment_property_id` on property_id
- `idx_comment_user_id` on user_id
- `idx_comment_created_at` on created_at

### BuyingBoxCriteria
**Purpose**: User-defined investment criteria for property evaluation

**Fields**:
- `id` (UUID, Primary Key): Unique identifier
- `user_id` (UUID, Foreign Key): Owner reference
- `name` (String, NOT NULL): Criteria set name
- `min_cap_rate` (Decimal(5,2)): Minimum cap rate
- `min_cash_on_cash` (Decimal(5,2)): Minimum CoC return
- `max_purchase_price` (Decimal(12,2)): Maximum purchase price
- `min_rent_to_value` (Decimal(5,2)): Minimum RTV ratio
- `max_year_built` (Integer): Maximum acceptable year built
- `min_year_built` (Integer): Minimum acceptable year built
- `created_at` (Timestamp): Criteria creation time
- `updated_at` (Timestamp): Last modification time
- `is_active` (Boolean, Default: true): Whether criteria is active

**JSON Fields**:
- `location_preferences` (JSON): Preferred areas, exclusions
- `property_type_preferences` (JSON): Single family, multi-family, etc.

**Validation Rules**:
- Name required, max 100 characters
- Numeric criteria must be positive where applicable
- Year range must be logical (min <= max)

**Indexes**:
- `idx_criteria_user_id` on user_id
- `idx_criteria_active` on is_active

## Calculation Formulas

### Net Operating Income (NOI)
```
NOI = (Monthly Rent × 12) - Annual Operating Expenses
Annual Operating Expenses = Insurance + Property Taxes + HOA + (Monthly Rent × 12 × Vacancy Rate) + Maintenance + Management + Utilities
```

### Cap Rate
```
Cap Rate = (NOI / Purchase Price) × 100
```

### Cash-on-Cash Return
```
Annual Cash Flow = NOI - Annual Debt Service
Cash-on-Cash Return = (Annual Cash Flow / Initial Cash Investment) × 100
Initial Cash Investment = Down Payment + Closing Costs
```

### Monthly Mortgage Payment
```
P = L[c(1 + c)^n]/[(1 + c)^n - 1]
Where: P = payment, L = loan amount, c = monthly interest rate, n = number of payments
```

### Rent-to-Value Ratio (RTV)
```
RTV = (Monthly Rent × 12 / Purchase Price) × 100
```

### Gross Rent Multiplier (GRM)
```
GRM = Purchase Price / (Monthly Rent × 12)
```

## Database Schema Migration Strategy

### Phase 1: Core Tables
1. Users table with authentication
2. Properties table with basic fields
3. Comments table for collaboration

### Phase 2: Financial Data
1. PropertyValuation table
2. FinancialMetrics table
3. BuyingBoxCriteria table

### Phase 3: Optimization
1. Add indexes for performance
2. Add triggers for metric recalculation
3. Add constraints and validations

## Data Integrity Rules

### Foreign Key Constraints
- All foreign keys use CASCADE DELETE for data consistency
- Property deletion removes all associated comments, valuations, and metrics
- User deletion transfers properties to system user or requires reassignment

### Business Logic Constraints
- Property must have purchase price to calculate metrics
- Financial metrics automatically recalculated when property data changes
- Valuations must be from approved sources only

### Data Validation
- All monetary values stored in cents to avoid floating point issues
- Dates validated to be reasonable (not in future for historical data)
- Percentages stored as decimals (0.05 for 5%)

---

*Data model designed for PostgreSQL with GORM ORM, optimized for the specific queries needed by rental property analysis workflows.*
