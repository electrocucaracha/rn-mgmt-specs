// API client for communicating with the backend
class ApiClient {
  constructor() {
    this.baseURL = '/api/v1';
    this.token = localStorage.getItem('auth_token');
  }

  // Set authentication token
  setToken(token) {
    this.token = token;
    if (token) {
      localStorage.setItem('auth_token', token);
    } else {
      localStorage.removeItem('auth_token');
    }
  }

  // Get authentication token
  getToken() {
    return this.token;
  }

  // Make HTTP request
  async request(endpoint, options = {}) {
    const url = `${this.baseURL}${endpoint}`;
    
    const config = {
      headers: {
        'Content-Type': 'application/json',
        ...options.headers,
      },
      ...options,
    };

    // Add auth token if available
    if (this.token) {
      config.headers.Authorization = `Bearer ${this.token}`;
    }

    try {
      const response = await fetch(url, config);
      const data = await response.json();

      if (!response.ok) {
        throw new Error(data.error || `HTTP error! status: ${response.status}`);
      }

      return data;
    } catch (error) {
      console.error('API request failed:', error);
      throw error;
    }
  }

  // Authentication methods
  async register(userData) {
    return this.request('/auth/register', {
      method: 'POST',
      body: JSON.stringify(userData),
    });
  }

  async login(credentials) {
    const response = await this.request('/auth/login', {
      method: 'POST',
      body: JSON.stringify(credentials),
    });
    
    if (response.token) {
      this.setToken(response.token);
    }
    
    return response;
  }

  async logout() {
    try {
      await this.request('/auth/logout', {
        method: 'POST',
      });
    } finally {
      this.setToken(null);
    }
  }

  // Property methods
  async getProperties(params = {}) {
    const queryString = new URLSearchParams(params).toString();
    const endpoint = `/properties${queryString ? `?${queryString}` : ''}`;
    return this.request(endpoint);
  }

  async getProperty(id) {
    return this.request(`/properties/${id}`);
  }

  async createProperty(propertyData) {
    return this.request('/properties', {
      method: 'POST',
      body: JSON.stringify(propertyData),
    });
  }

  async updateProperty(id, propertyData) {
    return this.request(`/properties/${id}`, {
      method: 'PUT',
      body: JSON.stringify(propertyData),
    });
  }

  async deleteProperty(id) {
    return this.request(`/properties/${id}`, {
      method: 'DELETE',
    });
  }

  // Property metrics methods
  async getPropertyMetrics(propertyId) {
    return this.request(`/properties/${propertyId}/metrics`);
  }

  async recalculateMetrics(propertyId) {
    return this.request(`/properties/${propertyId}/metrics`, {
      method: 'POST',
    });
  }

  // Property valuations methods
  async getPropertyValuations(propertyId) {
    return this.request(`/properties/${propertyId}/valuations`);
  }

  async addPropertyValuation(propertyId, valuationData) {
    return this.request(`/properties/${propertyId}/valuations`, {
      method: 'POST',
      body: JSON.stringify(valuationData),
    });
  }

  // Comments methods
  async getPropertyComments(propertyId) {
    return this.request(`/properties/${propertyId}/comments`);
  }

  async addPropertyComment(propertyId, commentData) {
    return this.request(`/properties/${propertyId}/comments`, {
      method: 'POST',
      body: JSON.stringify(commentData),
    });
  }

  async updateComment(id, commentData) {
    return this.request(`/comments/${id}`, {
      method: 'PUT',
      body: JSON.stringify(commentData),
    });
  }

  async deleteComment(id) {
    return this.request(`/comments/${id}`, {
      method: 'DELETE',
    });
  }

  // Buying criteria methods
  async getBuyingCriteria() {
    return this.request('/buying-criteria');
  }

  async createBuyingCriteria(criteriaData) {
    return this.request('/buying-criteria', {
      method: 'POST',
      body: JSON.stringify(criteriaData),
    });
  }

  async updateBuyingCriteria(id, criteriaData) {
    return this.request(`/buying-criteria/${id}`, {
      method: 'PUT',
      body: JSON.stringify(criteriaData),
    });
  }

  async deleteBuyingCriteria(id) {
    return this.request(`/buying-criteria/${id}`, {
      method: 'DELETE',
    });
  }

  async compareProperties(comparisonData) {
    return this.request('/properties/compare', {
      method: 'POST',
      body: JSON.stringify(comparisonData),
    });
  }
}

// Create and export a singleton instance
const apiClient = new ApiClient();
export default apiClient;
