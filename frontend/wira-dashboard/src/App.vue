<template>
  <div id="app">
    <AppDashboard />  <!-- Add this line to use the component -->
  </div>
</template>

<script>
// Import the AppDashboard component
import AppDashboard from './components/AppDashboard.vue';  

export default {
  name: 'App',
  components: {
    AppDashboard,  // Register it as a component
  },
};
</script>
