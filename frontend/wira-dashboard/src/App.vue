<template>
  <div id="app">
    <!-- Conditionally render the LoginPage or AppDashboard based on isLoggedIn -->
    <LoginPage v-if="!isLoggedIn" @login-success="handleLoginSuccess" />
    <AppDashboard v-else />
  </div>
</template>

<script>
// Import the components
import AppDashboard from './components/AppDashboard.vue';
import LoginPage from './components/LoginPage.vue';


export default {
  name: 'App',
  components: {
    AppDashboard,
    LoginPage
  },
  data() {
    return {
      isLoggedIn: false // Track the login status
    };
  },
  methods: {
    // Handle successful login and update the login status
    handleLoginSuccess() {
      this.isLoggedIn = true;
    }
  }
};
</script>
