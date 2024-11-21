<template>
  <div class="container mx-auto p-4">
    <h1 class="text-2xl font-bold mb-4">WIRA Ranking Dashboard</h1>
    <div class="flex justify-between mb-4">
      <input
        v-model="searchQuery"
        type="text"
        placeholder="Search players"
        class="px-4 py-2 border rounded"
      />
      <button @click="fetchData" class="px-4 py-2 bg-blue-500 text-white rounded">
        Search
      </button>
    </div>
    <div class="overflow-x-auto">
      <table class="min-w-full bg-white border border-gray-200 rounded-lg">
        <thead>
          <tr>
            <th class="px-4 py-2 border">Rank</th>
            <th class="px-4 py-2 border">Player</th>
            <th class="px-4 py-2 border">Score</th>
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
    <div class="mt-4">
      <button @click="paginate('prev')" :disabled="page === 1" class="px-4 py-2 bg-gray-500 text-white rounded">
        Previous
      </button>
      <span class="mx-2">{{ page }}</span>
      <button @click="paginate('next')" :disabled="page === totalPages" class="px-4 py-2 bg-gray-500 text-white rounded">
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
    };
  },
  methods: {
    async fetchData() {
      const response = await fetch(
        `http://localhost:8080/players?page=${this.page}&search=${this.searchQuery}`
      );
      const data = await response.json();
      console.log(data);
      this.players = data.players;
      this.totalPages = data.totalPages;
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
  mounted() {
    this.fetchData();
  },
};
</script>

<style scoped>
/* You can add custom styles here */
</style>
