# Deployment Guide: Rental Property Management Platform

## 🚀 Production Deployment

### Prerequisites
- Docker or Podman
- Docker Compose or Podman Compose
- 2GB+ RAM available
- 5GB+ disk space

### Quick Production Deployment

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd rental-property-management
   ```

2. **Configure environment**
   ```bash
   cp backend/.env.example backend/.env
   # Edit backend/.env with production values
   ```

3. **Start production environment**
   ```bash
   make prod
   # OR manually:
   docker-compose up -d
   ```

4. **Verify deployment**
   ```bash
   # Check all services are running
   docker-compose ps
   
   # Check health
   curl http://localhost:8080/api/v1/health
   
   # Access application
   open http://localhost
   ```

### Environment Configuration

#### Backend Environment Variables (.env)
```bash
# Database
DB_HOST=postgres
DB_PORT=5432
DB_USER=rental_user
DB_PASSWORD=change-this-in-production
DB_NAME=rental_property_mgmt
DB_SSLMODE=require

# Server
PORT=8080
ENV=production

# JWT Authentication
JWT_SECRET=your-super-secret-jwt-key-min-32-chars
JWT_EXPIRES_IN=24h

# CORS
CORS_ORIGINS=https://yourdomain.com
```

#### Production Security Checklist
- [ ] Change default database password
- [ ] Set strong JWT secret (32+ characters)
- [ ] Configure SSL/TLS certificates
- [ ] Set secure CORS origins
- [ ] Enable PostgreSQL SSL mode
- [ ] Configure firewall rules
- [ ] Set up monitoring and logging

## 🌐 Cloud Deployment Options

### AWS Deployment with ECS

1. **Build and push images**
   ```bash
   # Build images
   docker build -t rental-backend ./backend
   docker build -t rental-frontend ./frontend
   
   # Tag for ECR
   docker tag rental-backend:latest 123456789.dkr.ecr.region.amazonaws.com/rental-backend:latest
   docker tag rental-frontend:latest 123456789.dkr.ecr.region.amazonaws.com/rental-frontend:latest
   
   # Push to ECR
   docker push 123456789.dkr.ecr.region.amazonaws.com/rental-backend:latest
   docker push 123456789.dkr.ecr.region.amazonaws.com/rental-frontend:latest
   ```

2. **Deploy with ECS Service**
   - Use provided `ecs-task-definition.json`
   - Configure Application Load Balancer
   - Set up RDS PostgreSQL instance
   - Configure ECS Service with auto-scaling

### Google Cloud Run Deployment

1. **Build and deploy backend**
   ```bash
   cd backend
   gcloud builds submit --tag gcr.io/PROJECT-ID/rental-backend
   gcloud run deploy rental-backend \
     --image gcr.io/PROJECT-ID/rental-backend \
     --platform managed \
     --region us-central1 \
     --allow-unauthenticated
   ```

2. **Build and deploy frontend**
   ```bash
   cd frontend
   gcloud builds submit --tag gcr.io/PROJECT-ID/rental-frontend
   gcloud run deploy rental-frontend \
     --image gcr.io/PROJECT-ID/rental-frontend \
     --platform managed \
     --region us-central1 \
     --allow-unauthenticated
   ```

### Digital Ocean App Platform

1. **Create app.yaml**
   ```yaml
   name: rental-property-mgmt
   services:
   - name: backend
     source_dir: /backend
     github:
       repo: your-username/rental-property-mgmt
       branch: main
     run_command: ./main
     environment_slug: go
     instance_count: 1
     instance_size_slug: basic-xxs
     envs:
     - key: DB_HOST
       value: ${db.HOSTNAME}
     - key: DB_PASSWORD
       value: ${db.PASSWORD}
   
   - name: frontend
     source_dir: /frontend
     github:
       repo: your-username/rental-property-mgmt
       branch: main
     run_command: nginx -g 'daemon off;'
     environment_slug: node-js
     instance_count: 1
     instance_size_slug: basic-xxs
   
   databases:
   - name: db
     engine: PG
     version: "14"
   ```

## 🔧 Development Deployment

### Local Development with Hot Reload

1. **Start development environment**
   ```bash
   make dev
   ```
   
   This starts:
   - PostgreSQL database (port 5432)
   - Backend with Air hot reload (port 8080)
   - Frontend with Vite dev server (port 5173)

2. **Development workflow**
   ```bash
   # View logs
   make logs-dev
   
   # Run tests
   make test
   
   # Reset database
   make db-reset
   
   # Stop everything
   make stop
   ```

### Manual Local Development

1. **Start PostgreSQL**
   ```bash
   # Using Docker
   docker run -d \
     --name rental-postgres \
     -e POSTGRES_DB=rental_property_mgmt \
     -e POSTGRES_USER=rental_user \
     -e POSTGRES_PASSWORD=rental_password \
     -p 5432:5432 \
     postgres:14
   ```

2. **Start backend**
   ```bash
   cd backend
   cp .env.example .env
   go mod download
   go run cmd/server/main.go
   ```

3. **Start frontend**
   ```bash
   cd frontend
   npm install
   npm run dev
   ```

## 📊 Monitoring and Maintenance

### Health Checks
- Backend health: `GET /api/v1/health`
- Database connectivity: Built into backend health check
- Frontend: HTTP 200 on any static file request

### Logs
```bash
# View all logs
make logs

# Follow specific service logs
docker-compose logs -f backend
docker-compose logs -f postgres

# Search logs
docker-compose logs backend | grep ERROR
```

### Database Maintenance
```bash
# Backup database
docker-compose exec postgres pg_dump -U rental_user rental_property_mgmt > backup.sql

# Restore database
docker-compose exec -i postgres psql -U rental_user rental_property_mgmt < backup.sql

# Database shell access
docker-compose exec postgres psql -U rental_user rental_property_mgmt
```

### Updates and Migrations
```bash
# Pull latest changes
git pull origin main

# Rebuild and restart
make stop
make build
make prod

# Check migration status
docker-compose exec backend ./main --migrate-status
```

## 🔒 Security Considerations

### Production Security
1. **Database Security**
   - Use strong passwords
   - Enable SSL connections
   - Restrict network access
   - Regular backups

2. **Application Security**
   - Secure JWT secrets
   - HTTPS-only in production
   - Proper CORS configuration
   - Input validation and sanitization

3. **Infrastructure Security**
   - Firewall configuration
   - Regular security updates
   - Container image scanning
   - Access logging

### SSL/TLS Setup
```nginx
# Add to nginx.conf for HTTPS
server {
    listen 443 ssl http2;
    ssl_certificate /etc/ssl/certs/cert.pem;
    ssl_certificate_key /etc/ssl/private/key.pem;
    
    # Security headers
    add_header Strict-Transport-Security "max-age=31536000; includeSubDomains" always;
    add_header X-Frame-Options "SAMEORIGIN" always;
    add_header X-Content-Type-Options "nosniff" always;
}
```

## 📈 Scaling

### Horizontal Scaling
```bash
# Scale backend instances
docker-compose up -d --scale backend=3

# Load balancer configuration needed for multiple instances
```

### Performance Optimization
- Database connection pooling (implemented)
- Redis caching layer (future enhancement)
- CDN for static assets
- Database query optimization
- Container resource limits

## 🆘 Disaster Recovery

### Backup Strategy
1. **Database Backups**
   - Daily automated backups
   - Point-in-time recovery capability
   - Off-site backup storage

2. **Application Backups**
   - Configuration files
   - Environment variables
   - SSL certificates

3. **Recovery Procedures**
   - Documented rollback process
   - Tested restore procedures
   - Monitoring and alerting

---

For additional support, refer to the main [README.md](README.md) or raise an issue in the repository.
