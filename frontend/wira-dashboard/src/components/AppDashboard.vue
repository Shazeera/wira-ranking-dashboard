<template>
  <div class="dashboard-container">
    <h1 class="dashboard-title">WIRA Ranking Dashboard</h1>

    <!-- Search bar and button -->
    <div class="search-container">
      <div class="search-box">
        <input
          v-model="searchQuery"
          type="text"
          placeholder="Search players"
          class="search-input"
        />
        <button @click="fetchData" class="search-btn">Search</button>
      </div>
    </div>

    <!-- Loading Progress Bar -->
    <div v-if="loading" class="progress-container">
      <div class="progress-bar"></div>
    </div>

    <!-- Players Table -->
    <div v-else class="table-container">
      <table class="players-table">
        <thead>
          <tr>
            <th class="table-header">Rank</th>
            <th class="table-header">Player Username</th>
            <th class="table-header">Score</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(player, index) in players" :key="index">
            <td class="table-cell">{{ player.class_id }}</td>
            <td class="table-cell">{{ player.username }}</td>
            <td class="table-cell">{{ player.reward_score }}</td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- Pagination Controls -->
    <div class="pagination-container">
      <button 
        @click="paginate('prev')" 
        :disabled="page === 1" 
        class="pagination-btn"
      >
        Previous
      </button>
      <span class="pagination-text">Page {{ page }} of {{ totalPages }}</span>
      <button 
        @click="paginate('next')" 
        :disabled="page === totalPages" 
        class="pagination-btn"
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
      loading: false,
      backendUrl: process.env.VUE_APP_BACKEND_URL || 'http://172.21.48.1:8080',
    };
  },
  methods: {
    async fetchData() {
      this.loading = true;
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
        this.loading = false;
      }
    },
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
    searchQuery() {
      this.page = 1;
      this.fetchData();
    },
  },
  mounted() {
    this.fetchData();
  },
};
</script>

<style scoped>
/* General Styling */
.dashboard-container {
  font-family: 'Roboto', sans-serif;
  padding: 20px;
  max-width: 1200px;
  margin: 0 auto;
}

.dashboard-title {
  text-align: center;
  font-size: 2.5rem;
  font-weight: bold;
  margin-bottom: 20px;
  color: #333;
}

/* Search Bar */
.search-container {
  display: flex;
  justify-content: center;
  margin-bottom: 20px;
}

.search-box {
  display: flex;
  width: 100%;
  max-width: 600px;
  gap: 10px;
}

.search-input {
  flex: 1;
  padding: 12px;
  font-size: 1rem;
  border: 1px solid #ccc;
  border-radius: 5px;
  outline: none;
  transition: border-color 0.3s;
}

.search-input:focus {
  border-color: #333;
}

.search-btn {
  padding: 12px 20px;
  background-color: #333;
  color: white;
  font-size: 1rem;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  transition: background-color 0.3s;
}

.search-btn:hover {
  background-color: #555;
}

/* Progress Bar */
.progress-container {
  width: 100%;
  height: 5px;
  background: #f3f3f3;
  overflow: hidden;
}

.progress-bar {
  width: 100%;
  height: 100%;
  background: #333;
  animation: progress-animation 1.5s infinite;
}

@keyframes progress-animation {
  0% {
    transform: translateX(-100%);
  }
  100% {
    transform: translateX(100%);
  }
}

/* Table */
.table-container {
  margin-top: 20px;
  overflow-x: auto;
}

.players-table {
  width: 100%;
  border-collapse: collapse;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

.table-header {
  background-color: #333;
  color: white;
  padding: 12px;
  font-weight: bold;
}

.table-cell {
  padding: 12px;
  border: 1px solid #ddd;
  text-align: left;
}

/* Pagination */
.pagination-container {
  display: flex;
  justify-content: center;
  align-items: center;
  margin-top: 20px;
  gap: 20px;
}

.pagination-btn {
  padding: 10px 20px;
  background-color: #333;
  color: white;
  font-size: 1rem;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  transition: background-color 0.3s;
}

.pagination-btn:hover {
  background-color: #555;
}

.pagination-btn:disabled {
  background-color: #ccc;
  cursor: not-allowed;
}

.pagination-text {
  font-size: 1.2rem;
  color: #333;
}
</style>

