<template>
  <div class="container mx-auto p-4">
    <h1 class="text-2xl font-bold mb-4 text-center md:text-left">WIRA Ranking Dashboard</h1>
    
    <!-- Search bar and button -->
    <div class="flex flex-col md:flex-row justify-between mb-4">
      <div class="mb-2 md:mb-0 w-full md:w-auto">
        <input
          v-model="searchQuery"
          type="text"
          placeholder="Search players"
          class="w-full md:w-64 px-4 py-2 border rounded"
        />
      </div>
      <button 
        @click="fetchData" 
        class="w-full md:w-auto px-4 py-2 bg-blue-500 text-white rounded mt-2 md:mt-0"
      >
        Search
      </button>
    </div>
    
    <!-- Loading Spinner -->
    <div v-if="loading" class="text-center py-4">
      <span>Loading...</span>
    </div>

    <!-- Players Table -->
    <div v-else class="overflow-x-auto">
      <table class="min-w-full bg-white border border-gray-200 rounded-lg">
        <thead>
          <tr>
            <th class="px-4 py-2 border text-left">Rank</th>
            <th class="px-4 py-2 border text-left">Player Username</th>
            <th class="px-4 py-2 border text-left">Score</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(player, index) in players" :key="index">
            <td class="px-4 py-2 border">{{ player.class_id }}</td>
            <td class="px-4 py-2 border">{{ player.username }}</td>
            <td class="px-4 py-2 border">{{ player.reward_score }}</td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- Pagination Controls -->
    <div class="mt-4 flex justify-center md:justify-start items-center">
      <button 
        @click="paginate('prev')" 
        :disabled="page === 1" 
        class="px-4 py-2 bg-gray-500 text-white rounded mr-2"
      >
        Previous
      </button>
      <span class="mx-2">{{ page }}</span>
      <button 
        @click="paginate('next')" 
        :disabled="page === totalPages" 
        class="px-4 py-2 bg-gray-500 text-white rounded ml-2"
      >
        Next
      </button>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      players: [],
      searchQuery: '',
      page: 1,
      totalPages: 1,
      loading: false, // Loading state
      // Configurable backend URL for development and production
      backendUrl: process.env.VUE_APP_BACKEND_URL || 'http://172.21.48.1:8080', //  backend address
    };
  },
  methods: {
    // Fetch data from the backend with error handling
    async fetchData() {
      this.loading = true; // Set loading state
      try {
        const response = await fetch(
          `${this.backendUrl}/players?page=${this.page}&search=${this.searchQuery}`
        );
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }
        const data = await response.json();
        this.players = data.players;
        this.totalPages = data.totalPages;
      } catch (error) {
        console.error('Fetch error:', error);
        alert('Error fetching player data. Please try again later.');
      } finally {
        this.loading = false; // Reset loading state after the request
      }
    },
    // Pagination function
    paginate(direction) {
      if (direction === 'next' && this.page < this.totalPages) {
        this.page++;
      } else if (direction === 'prev' && this.page > 1) {
        this.page--;
      }
      this.fetchData();
    },
  },
  watch: {
    // Reset page to 1 when the search query changes
    searchQuery() {
      this.page = 1;
      this.fetchData();
    }
  },
  mounted() {
    // Fetch data on component mount
    this.fetchData();
  },
};
</script>

<style scoped>
/* Add any custom styles here if needed */
</style>
