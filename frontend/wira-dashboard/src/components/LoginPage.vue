<template>
  <div class="login-container">
    <div class="login-box">
      <h1>Welcome to WIRA Application</h1>
      <form @submit.prevent="login">
        <div class="input-group">
          <label for="username">Username</label>
          <input type="text" v-model="username" id="username" placeholder="Enter your username" required />
        </div>
        <div class="input-group">
          <label for="password">Password</label>
          <input type="password" v-model="password" id="password" placeholder="Enter your password" required />
        </div>
        <div class="input-group">
          <label for="twofa">2FA Code</label>
          <input type="text" v-model="twoFACode" id="twofa" placeholder="Enter 2FA code" required />
        </div>
        <button type="submit" class="login-btn">Login</button>
        <div v-if="errorMessage" class="error">{{ errorMessage }}</div>
      </form>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      username: '',
      password: '',
      twofaCode: '',
      errorMessage: ''
    };
  },
  methods: {
    async login() {
      try {
        const response = await fetch('http://172.21.48.1:8080/login', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({
            username: this.username,
            password: this.password,
            two_fa_code: this.twoFACode
          })
        });

        if (response.ok) {
          const data = await response.json();
          const authToken = data.token; // Assume the token is in the response body

          // Store the auth token in localStorage
          localStorage.setItem('authToken', authToken);

          // Emit the login-success event to notify the parent component
          this.$emit('login-success');
        } else {
          // If the login fails, show an error message
          const errorData = await response.json();
          this.errorMessage = errorData.error || 'Login failed. Please try again.';
        }
      } catch (error) {
        console.error('Login error:', error);
        this.errorMessage = 'An error occurred. Please try again later.';
      }
    }
  }
};
</script>

<style scoped>
/* Apply Arial font to all text */
* {
  font-family: Arial, sans-serif;
}

/* Main container styling */
.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background-color: #f2f4f7;
}

/* Styling the login box */
.login-box {
  background-color: #fff;
  padding: 30px;
  border-radius: 8px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  width: 100%;
  max-width: 400px;
}

/* Heading */
h1 {
  text-align: center;
  margin-bottom: 20px;
  font-size: 2rem;
  color: #333;
}

/* Input group styling */
.input-group {
  margin-bottom: 20px;
}

.input-group label {
  display: block;
  font-weight: bold;
  margin-bottom: 8px;
  color: #555;
}

.input-group input {
  width: 100%;
  padding: 12px;
  font-size: 1rem;
  border-radius: 4px;
  border: 1px solid #ccc;
  box-sizing: border-box;
}

.input-group input:focus {
  border-color: #4CAF50;
  outline: none;
}

/* Submit button styling */
.login-btn {
  width: 100%;
  padding: 12px;
  background-color: #333; /* Dark gray */
  color: white;
  border: none;
  border-radius: 4px;
  font-size: 1rem;
  cursor: pointer;
  transition: background-color 0.3s, color 0.3s;
}

.login-btn:hover {
  background-color: black; /* Black on hover */
  color: #ddd; /* Light gray text on hover */
}

/* Error message styling */
.error {
  color: red;
  font-size: 0.9rem;
  margin-top: 10px;
  text-align: center;
}
</style>
