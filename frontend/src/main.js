import apiClient from './services/api.js';

class RentalPropertyApp {
  constructor() {
    this.currentUser = null;
    this.currentSection = 'login';
    this.init();
  }

  init() {
    this.setupEventListeners();
    this.checkAuthStatus();
    this.showSection('login');
  }

  setupEventListeners() {
    // Navigation
    document.getElementById('nav-login').addEventListener('click', () => this.showSection('login'));
    document.getElementById('nav-register').addEventListener('click', () => this.showSection('register'));
    document.getElementById('nav-properties').addEventListener('click', () => this.showSection('properties'));
    document.getElementById('nav-add-property').addEventListener('click', () => this.showSection('add-property'));
    document.getElementById('nav-criteria').addEventListener('click', () => this.showSection('criteria'));

    // Forms
    document.getElementById('login-form').addEventListener('submit', (e) => this.handleLogin(e));
    document.getElementById('register-form').addEventListener('submit', (e) => this.handleRegister(e));
    document.getElementById('property-form').addEventListener('submit', (e) => this.handleAddProperty(e));
    
    // Logout
    document.getElementById('logout-btn').addEventListener('click', () => this.handleLogout());
  }

  showSection(sectionName) {
    // Hide all sections
    const sections = document.querySelectorAll('.section');
    sections.forEach(section => section.classList.add('hidden'));

    // Show target section
    document.getElementById(`${sectionName}-section`).classList.remove('hidden');

    // Update navigation
    const navButtons = document.querySelectorAll('.nav-btn');
    navButtons.forEach(btn => btn.classList.remove('active'));
    document.getElementById(`nav-${sectionName}`).classList.add('active');

    this.currentSection = sectionName;

    // Load data for certain sections
    if (sectionName === 'properties' && this.currentUser) {
      this.loadProperties();
    } else if (sectionName === 'criteria' && this.currentUser) {
      this.loadBuyingCriteria();
    }
  }

  checkAuthStatus() {
    const token = apiClient.getToken();
    if (token) {
      // TODO: Validate token with backend
      this.updateAuthUI(true);
    } else {
      this.updateAuthUI(false);
    }
  }

  updateAuthUI(isAuthenticated) {
    const userInfo = document.getElementById('user-info');
    const logoutBtn = document.getElementById('logout-btn');
    const authNavButtons = ['nav-properties', 'nav-add-property', 'nav-criteria'];
    const guestNavButtons = ['nav-login', 'nav-register'];

    if (isAuthenticated) {
      userInfo.classList.remove('hidden');
      logoutBtn.classList.remove('hidden');
      
      authNavButtons.forEach(id => {
        document.getElementById(id).classList.remove('hidden');
      });
      
      guestNavButtons.forEach(id => {
        document.getElementById(id).classList.add('hidden');
      });

      if (this.currentSection === 'login' || this.currentSection === 'register') {
        this.showSection('properties');
      }
    } else {
      userInfo.classList.add('hidden');
      logoutBtn.classList.add('hidden');
      
      authNavButtons.forEach(id => {
        document.getElementById(id).classList.add('hidden');
      });
      
      guestNavButtons.forEach(id => {
        document.getElementById(id).classList.remove('hidden');
      });

      this.showSection('login');
    }
  }

  async handleLogin(e) {
    e.preventDefault();
    
    const email = document.getElementById('login-email').value;
    const password = document.getElementById('login-password').value;
    const errorElement = document.getElementById('login-error');

    try {
      const response = await apiClient.login({ email, password });
      this.currentUser = response.user;
      
      document.getElementById('user-info').textContent = `Welcome, ${response.user.first_name}!`;
      this.updateAuthUI(true);
      this.clearError('login-error');
      
    } catch (error) {
      this.showError('login-error', error.message);
    }
  }

  async handleRegister(e) {
    e.preventDefault();
    
    const userData = {
      email: document.getElementById('register-email').value,
      password: document.getElementById('register-password').value,
      first_name: document.getElementById('register-first-name').value,
      last_name: document.getElementById('register-last-name').value,
    };

    try {
      await apiClient.register(userData);
      this.showSuccess('register-success', 'Registration successful! Please login.');
      this.clearError('register-error');
      
      // Clear form
      document.getElementById('register-form').reset();
      
    } catch (error) {
      this.showError('register-error', error.message);
    }
  }

  async handleAddProperty(e) {
    e.preventDefault();
    
    const propertyData = {
      address: document.getElementById('property-address').value,
      purchase_price: parseFloat(document.getElementById('property-price').value),
      intended_rent: parseFloat(document.getElementById('property-rent').value) || null,
      year_built: parseInt(document.getElementById('property-year').value) || null,
    };

    try {
      await apiClient.createProperty(propertyData);
      this.showSuccess('property-success', 'Property added successfully!');
      this.clearError('property-error');
      
      // Clear form
      document.getElementById('property-form').reset();
      
      // Refresh properties list if we're on that section
      if (this.currentSection === 'properties') {
        this.loadProperties();
      }
      
    } catch (error) {
      this.showError('property-error', error.message);
    }
  }

  async handleLogout() {
    try {
      await apiClient.logout();
    } catch (error) {
      console.error('Logout error:', error);
    } finally {
      this.currentUser = null;
      this.updateAuthUI(false);
    }
  }

  async loadProperties() {
    try {
      const properties = await apiClient.getProperties();
      this.displayProperties(properties);
    } catch (error) {
      console.error('Failed to load properties:', error);
    }
  }

  displayProperties(properties) {
    const container = document.getElementById('properties-list');
    
    if (!properties || properties.length === 0) {
      container.innerHTML = '<p>No properties found. <a href="#" onclick="app.showSection(\'add-property\')">Add your first property</a>.</p>';
      return;
    }

    const html = properties.map(property => `
      <div style="border: 1px solid #ddd; padding: 15px; margin-bottom: 10px; border-radius: 4px;">
        <h3>${property.address}</h3>
        <p><strong>Purchase Price:</strong> $${property.purchase_price.toLocaleString()}</p>
        ${property.intended_rent ? `<p><strong>Intended Rent:</strong> $${property.intended_rent.toLocaleString()}/month</p>` : ''}
        ${property.year_built ? `<p><strong>Year Built:</strong> ${property.year_built}</p>` : ''}
        <p><strong>Added:</strong> ${new Date(property.created_at).toLocaleDateString()}</p>
      </div>
    `).join('');

    container.innerHTML = html;
  }

  async loadBuyingCriteria() {
    try {
      const criteria = await apiClient.getBuyingCriteria();
      this.displayBuyingCriteria(criteria);
    } catch (error) {
      console.error('Failed to load buying criteria:', error);
    }
  }

  displayBuyingCriteria(criteria) {
    const container = document.getElementById('criteria-list');
    
    if (!criteria || criteria.length === 0) {
      container.innerHTML = '<p>No buying criteria defined yet.</p>';
      return;
    }

    const html = criteria.map(criterion => `
      <div style="border: 1px solid #ddd; padding: 15px; margin-bottom: 10px; border-radius: 4px;">
        <h3>${criterion.name}</h3>
        ${criterion.min_cap_rate ? `<p><strong>Min Cap Rate:</strong> ${criterion.min_cap_rate}%</p>` : ''}
        ${criterion.min_cash_on_cash ? `<p><strong>Min Cash-on-Cash:</strong> ${criterion.min_cash_on_cash}%</p>` : ''}
        ${criterion.max_purchase_price ? `<p><strong>Max Purchase Price:</strong> $${criterion.max_purchase_price.toLocaleString()}</p>` : ''}
        <p><strong>Status:</strong> ${criterion.is_active ? 'Active' : 'Inactive'}</p>
      </div>
    `).join('');

    container.innerHTML = html;
  }

  showError(elementId, message) {
    const element = document.getElementById(elementId);
    element.textContent = message;
    element.classList.remove('hidden');
  }

  clearError(elementId) {
    const element = document.getElementById(elementId);
    element.classList.add('hidden');
  }

  showSuccess(elementId, message) {
    const element = document.getElementById(elementId);
    element.textContent = message;
    element.classList.remove('hidden');
    
    // Auto-hide success message after 3 seconds
    setTimeout(() => {
      element.classList.add('hidden');
    }, 3000);
  }
}

// Initialize the app
const app = new RentalPropertyApp();

// Make app globally available for debugging and HTML onclick handlers
window.app = app;
