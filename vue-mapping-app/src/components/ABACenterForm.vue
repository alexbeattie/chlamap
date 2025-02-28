<template>
  <div class="aba-center-form">
    <h1>ABA Center Entry Form</h1>
    
    <div v-if="successMessage" class="success-message">
      {{ successMessage }}
    </div>
    
    <div v-if="errorMessage" class="error-message">
      {{ errorMessage }}
    </div>
    
    <form @submit.prevent="submitForm">
      <div class="form-group">
        <label for="name">Center Name *</label>
        <input 
          type="text" 
          id="name" 
          v-model="formData.name" 
          required
        />
      </div>
      
      <div class="form-group">
        <label for="street">Street Address *</label>
        <input 
          type="text" 
          id="street" 
          v-model="formData.street" 
          required
        />
      </div>
      
      <div class="form-row">
        <div class="form-group">
          <label for="city">City *</label>
          <input 
            type="text" 
            id="city" 
            v-model="formData.city" 
            required
          />
        </div>
        
        <div class="form-group">
          <label for="zip">ZIP Code *</label>
          <input 
            type="text" 
            id="zip" 
            v-model="formData.zip" 
            pattern="[0-9]{5}"
            title="Five digit ZIP code"
            required
          />
        </div>
      </div>
      
      <div class="form-group">
        <label for="phone">Phone Number *</label>
        <input 
          type="tel" 
          id="phone" 
          v-model="formData.phone" 
          required
        />
      </div>
      
      <div class="form-group">
        <label for="serviceType">Service Type *</label>
        <select 
          id="serviceType" 
          v-model="formData.serviceType" 
          required
        >
          <option value="">Select service type</option>
          <option value="Home-Based">Home-Based</option>
          <option value="Center-Based">Center-Based</option>
          <option value="Both">Both Home and Center Based</option>
        </select>
      </div>
      
      <div class="form-group">
        <label for="waitlistAvailability">Waitlist Availability</label>
        <textarea 
          id="waitlistAvailability" 
          v-model="formData.waitlistAvailability" 
          rows="3"
        ></textarea>
      </div>
      
      <div class="form-group">
        <label for="waitlistNotes">Waitlist Notes</label>
        <textarea 
          id="waitlistNotes" 
          v-model="formData.waitlistNotes" 
          rows="2"
        ></textarea>
      </div>
      
      <div class="form-group">
        <label for="dxVerification">Diagnosis Verification Requirements</label>
        <textarea 
          id="dxVerification" 
          v-model="formData.dxVerification" 
          rows="2"
        ></textarea>
      </div>
      
      <div class="form-group">
        <label for="insuranceAccepted">Insurance Accepted (In Network)</label>
        <textarea 
          id="insuranceAccepted" 
          v-model="formData.insuranceAccepted" 
          rows="3"
        ></textarea>
      </div>
      
      <div class="form-group">
        <label for="mediCalPlans">Medi-Cal Managed Care Plans</label>
        <textarea 
          id="mediCalPlans" 
          v-model="formData.mediCalPlans" 
          rows="2"
        ></textarea>
      </div>
      
      <div class="form-group">
        <label for="notes">Additional Notes</label>
        <textarea 
          id="notes" 
          v-model="formData.notes" 
          rows="3"
        ></textarea>
      </div>
      
      <div class="form-actions">
        <button type="submit" :disabled="isSubmitting">
          {{ isSubmitting ? 'Submitting...' : 'Save Center' }}
        </button>
        <button type="button" @click="resetForm" :disabled="isSubmitting">
          Clear Form
        </button>
      </div>
    </form>
  </div>
</template>

<script>


export default {
  name: 'AbaCenterForm',
  data() {
    return {
      formData: {
        name: '',
        street: '',
        city: '',
        zip: '',
        phone: '',
        serviceType: '',
        waitlistAvailability: '',
        waitlistNotes: '',
        dxVerification: '',
        insuranceAccepted: '',
        mediCalPlans: '',
        notes: ''
      },
      isSubmitting: false,
      successMessage: '',
      errorMessage: ''
    };
  },
  methods: {
    async submitForm() {
      this.isSubmitting = true;
      this.successMessage = '';
      this.errorMessage = '';
      
      try {
        // Make API call to the backend
        const response = await fetch('http://localhost:3000/api/aba-centers', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(this.formData)
        });
        
        const data = await response.json();
        
        if (!response.ok) {
          throw new Error(data.error || 'Failed to save ABA center information');
        }
        
        // Handle success
        this.successMessage = 'ABA Center information saved successfully!';
        this.resetForm();
      } catch (error) {
        // Handle error
        console.error('Error submitting form:', error);
        this.errorMessage = error.message || 'An error occurred while saving the center information. Please try again.';
      } finally {
        this.isSubmitting = false;
      }
    },
    
    resetForm() {
      // Reset all form fields to their initial state
      this.formData = {
        name: '',
        street: '',
        city: '',
        zip: '',
        phone: '',
        serviceType: '',
        waitlistAvailability: '',
        waitlistNotes: '',
        dxVerification: '',
        insuranceAccepted: '',
        mediCalPlans: '',
        notes: ''
      };
    }
  }
};
</script>

<style scoped>
.aba-center-form {
  max-width: 800px;
  margin: 0 auto;
  padding: 20px;
  font-family: Arial, sans-serif;
}

h1 {
  text-align: center;
  color: #2c3e50;
  margin-bottom: 30px;
}

.form-group {
  margin-bottom: 20px;
}

.form-row {
  display: flex;
  gap: 20px;
}

.form-row .form-group {
  flex: 1;
}

label {
  display: block;
  margin-bottom: 5px;
  font-weight: bold;
}

input, select, textarea {
  width: 100%;
  padding: 10px;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 16px;
}

textarea {
  resize: vertical;
}

.form-actions {
  display: flex;
  gap: 15px;
  margin-top: 30px;
}

button {
  padding: 12px 20px;
  background-color: #4CAF50;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 16px;
  flex: 1;
}

button[type="button"] {
  background-color: #f1f1f1;
  color: #333;
}

button:hover {
  opacity: 0.9;
}

button:disabled {
  background-color: #cccccc;
  cursor: not-allowed;
}

.success-message {
  background-color: #dff0d8;
  color: #3c763d;
  padding: 15px;
  margin-bottom: 20px;
  border-radius: 4px;
}

.error-message {
  background-color: #f2dede;
  color: #a94442;
  padding: 15px;
  margin-bottom: 20px;
  border-radius: 4px;
}
</style>