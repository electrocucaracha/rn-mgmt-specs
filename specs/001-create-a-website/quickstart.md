# Quickstart Guide: Rental Property Management Platform

## Development Environment Setup

### Prerequisites
- Go 1.21 or higher
- Node.js 18 or higher
- PostgreSQL 14 or higher
- Git

### 1. Clone and Setup Repository
```bash
git clone <repository-url>
cd rental-property-management
```

### 2. Backend Setup
```bash
cd backend

# Install Go dependencies
go mod download

# Set up environment variables
cp .env.example .env
# Edit .env with your database credentials

# Run database migrations
go run cmd/migrate/main.go up

# Start the API server
go run cmd/server/main.go
```

The backend will be available at `http://localhost:8080`

### 3. Frontend Setup
```bash
cd frontend

# Install frontend dependencies
npm install

# Start the development server
npm run dev
```

The frontend will be available at `http://localhost:5173`

### 4. Database Setup
```sql
-- Create database (run as postgres user)
CREATE DATABASE rental_property_mgmt;
CREATE USER rental_user WITH PASSWORD 'rental_password';
GRANT ALL PRIVILEGES ON DATABASE rental_property_mgmt TO rental_user;
```

## Quick Test Scenarios

### Scenario 1: User Registration and Login
**Test Story**: New user creates account and logs in

1. **Open the application** in browser at `http://localhost:5173`
2. **Click "Register"** on the homepage
3. **Fill registration form**:
   - Email: test@example.com
   - Password: testpassword123
   - First Name: John
   - Last Name: Doe
4. **Submit registration** - Should redirect to dashboard
5. **Logout** using menu
6. **Login again** with same credentials
7. **Verify** successful login and dashboard access

**Expected Result**: User can register, logout, and login successfully

### Scenario 2: Property Creation and Basic Analysis
**Test Story**: User creates property and views calculated metrics

1. **Login** as registered user
2. **Navigate to "Add Property"** page
3. **Fill property form**:
   - Address: 123 Main St, Anytown, ST 12345
   - Year Built: 2000
   - Land Area: 6000 sq ft
   - Building Area: 1500 sq ft
   - Purchase Price: $250,000
   - Intended Rent: $2,100

4. **Add Operating Expenses**:
   - Insurance: $1,200/year
   - Property Taxes: $3,600/year
   - HOA: $0/year

5. **Add Financing Terms**:
   - Interest Rate: 7.5%
   - Loan Term: 30 years
   - Down Payment: 20%
   - Closing Costs: $5,000

6. **Add Operating Assumptions**:
   - Vacancy Rate: 5%
   - Maintenance: 10% of rent
   - Management: 8% of rent
   - Utilities: $0 (tenant paid)

7. **Save property** and view details page

**Expected Results**:
- Property saves successfully
- Financial metrics automatically calculated:
  - Monthly Mortgage Payment: ~$1,398
  - Net Operating Income: ~$19,392
  - Cap Rate: ~7.76%
  - Cash-on-Cash Return: ~15.2%
  - Cash to Close: $55,000
  - RTV: 10.08%
  - GRM: 9.92

### Scenario 3: Third-Party Valuations
**Test Story**: User adds third-party valuation data

1. **Navigate to property details** page
2. **Click "Add Valuation"** button
3. **Add Zillow market valuation**:
   - Source: Zillow
   - Type: Market Value
   - Value: $275,000
   - Date: Current date

4. **Add Redfin rental estimate**:
   - Source: Redfin
   - Type: Rental Estimate
   - Value: $2,200
   - Date: Current date

5. **View updated property page**

**Expected Results**:
- Valuations display in property details
- Average market value shown: $262,500 (property + Zillow average)
- Rental estimates compared to intended rent

### Scenario 4: Buying Box Criteria and Comparison
**Test Story**: User defines investment criteria and compares properties

1. **Navigate to "Buying Criteria"** page
2. **Create new criteria set**:
   - Name: "Conservative Investment"
   - Min Cap Rate: 7%
   - Min Cash-on-Cash: 12%
   - Max Purchase Price: $300,000
   - Min RTV: 8%

3. **Save criteria**
4. **Navigate to property list**
5. **Select "Compare Properties"**
6. **Choose criteria set** and properties to compare
7. **View comparison results**

**Expected Results**:
- Criteria saves successfully
- Properties compared against criteria
- Match/no-match indicators displayed
- Comparison scores calculated

### Scenario 5: Team Collaboration
**Test Story**: Multiple users collaborate on property analysis

1. **Create second user account** (different email)
2. **User 1**: Share property with User 2 (via collaboration feature)
3. **User 2**: Login and view shared property
4. **User 2**: Add comment on property: "Great location, but check roof condition"
5. **User 1**: Reply to comment: "Roof was inspected last year, in good condition"
6. **Both users**: View comment thread

**Expected Results**:
- Property sharing works correctly
- Comments display with author names
- Comment threads maintain proper order
- Real-time or near-real-time updates

### Scenario 6: Property Editing and Metric Recalculation
**Test Story**: User updates property data and sees metrics update

1. **Navigate to existing property**
2. **Click "Edit Property"**
3. **Update intended rent** from $2,100 to $2,300
4. **Update interest rate** from 7.5% to 6.8%
5. **Save changes**
6. **View updated metrics**

**Expected Results**:
- Metrics automatically recalculate
- New Cap Rate: ~8.44%
- New Cash-on-Cash Return: ~18.1%
- All dependent metrics update correctly

## API Testing with curl

### Authentication Test
```bash
# Register new user
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "password": "testpass123",
    "first_name": "Test",
    "last_name": "User"
  }'

# Login
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "password": "testpass123"
  }'
```

### Property Management Test
```bash
# Create property (replace TOKEN with actual JWT)
curl -X POST http://localhost:8080/api/v1/properties \
  -H "Authorization: Bearer TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "address": "123 Test St, Test City, TS 12345",
    "purchase_price": 250000,
    "intended_rent": 2100,
    "operating_expenses": {
      "insurance": 1200,
      "property_taxes": 3600,
      "hoa": 0
    },
    "financing_terms": {
      "interest_rate": 7.5,
      "loan_term": 30,
      "down_payment_percent": 20,
      "closing_costs": 5000
    }
  }'

# Get all properties
curl -X GET http://localhost:8080/api/v1/properties \
  -H "Authorization: Bearer TOKEN"
```

## Performance Validation

### Load Testing Setup
```bash
# Install k6 (load testing tool)
npm install -g k6

# Run basic load test
k6 run load-tests/api-test.js
```

### Expected Performance
- **API Response Time**: < 500ms for standard queries
- **Property List Loading**: < 2 seconds for 100 properties
- **Metric Calculations**: < 100ms per property
- **Concurrent Users**: Support 100+ simultaneous users

## Troubleshooting

### Common Issues

**Backend won't start**:
- Check Go version: `go version`
- Verify database connection
- Check port 8080 availability

**Frontend won't start**:
- Check Node.js version: `node --version`
- Clear npm cache: `npm cache clean --force`
- Check port 5173 availability

**Database connection fails**:
- Verify PostgreSQL is running
- Check .env file configuration
- Ensure database exists and user has permissions

**Calculations seem wrong**:
- Verify all required fields are filled
- Check decimal precision in database
- Validate formula implementations against documentation

### Getting Help
- Check API documentation at `http://localhost:8080/docs`
- Review backend logs for error details
- Check browser console for frontend errors
- Verify all test scenarios pass before reporting issues

---

*This quickstart guide validates the core functionality described in the feature specification through hands-on testing scenarios.*
