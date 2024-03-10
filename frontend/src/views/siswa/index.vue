<script setup>
import Navbar from './../layout/mainView.vue'
</script>

<template>
  <Navbar/>

    <div class="mt-20 sm:mx-2 lg:mx-10 place-items-center">
      <div class="">
        <div class="text-end mb-5">
          <RouterLink to="/murid/create" class="bg-indigo-600 px-3 py-2 rounded-sm text-white text-sm">
            Tambah
          </RouterLink>
        </div>
        <div class="relative overflow-x-auto shadow-md sm:rounded-lg">
          <table class="w-full text-sm text-left rtl:text-right text-gray-500 dark:text-gray-400">
            <thead class="text-xs text-gray-700 bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
              <tr>
                <th class="px-6 py-3">No</th>
                <th class="px-6 py-3 text-center">Nama</th>
                <th class="px-6 py-3 text-center">Email</th>
                <th class="px-6 py-3 text-center">NIS</th>
                <th class="px-6 py-3 text-center">Class</th>
                <th class="px-6 py-3 text-center">Action</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="(item, Index) in items" :key="item.id" class="bg-white border-b dark:bg-gray-800 dark:border-gray-700">
                  <td class="px-6 py-4">{{ Index + 1 }}</td>
                  <td class="px-6 py-4 text-center">{{ item.name }}</td>
                  <td class="px-6 py-4 text-center">{{ item.email }}</td>
                  <td class="px-6 py-4 text-center">{{ item.nis }}</td>
                  <td class="px-6 py-4 text-center">{{ item.Class.name +' '+ item.Class.jurusan }}</td>
                  <td class="px-6 py-4 flex justify-center">
                      <RouterLink :to="'/murid/' + item.id" class="font-medium bg-blue-500 text-white p-2 rounded mx-1">
                        <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-pencil" viewBox="0 0 16 16">
                          <path d="M12.146.146a.5.5 0 0 1 .708 0l3 3a.5.5 0 0 1 0 .708l-10 10a.5.5 0 0 1-.168.11l-5 2a.5.5 0 0 1-.65-.65l2-5a.5.5 0 0 1 .11-.168zM11.207 2.5 13.5 4.793 14.793 3.5 12.5 1.207zm1.586 3L10.5 3.207 4 9.707V10h.5a.5.5 0 0 1 .5.5v.5h.5a.5.5 0 0 1 .5.5v.5h.293zm-9.761 5.175-.106.106-1.528 3.821 3.821-1.528.106-.106A.5.5 0 0 1 5 12.5V12h-.5a.5.5 0 0 1-.5-.5V11h-.5a.5.5 0 0 1-.468-.325"/>
                        </svg>
                      </RouterLink>
                      <Button @click="deleteItem(item.id)" class="font-medium bg-red-500 text-white p-2 rounded mx-1">
                        <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-trash3" viewBox="0 0 16 16">
                          <path d="M6.5 1h3a.5.5 0 0 1 .5.5v1H6v-1a.5.5 0 0 1 .5-.5M11 2.5v-1A1.5 1.5 0 0 0 9.5 0h-3A1.5 1.5 0 0 0 5 1.5v1H1.5a.5.5 0 0 0 0 1h.538l.853 10.66A2 2 0 0 0 4.885 16h6.23a2 2 0 0 0 1.994-1.84l.853-10.66h.538a.5.5 0 0 0 0-1zm1.958 1-.846 10.58a1 1 0 0 1-.997.92h-6.23a1 1 0 0 1-.997-.92L3.042 3.5zm-7.487 1a.5.5 0 0 1 .528.47l.5 8.5a.5.5 0 0 1-.998.06L5 5.03a.5.5 0 0 1 .47-.53Zm5.058 0a.5.5 0 0 1 .47.53l-.5 8.5a.5.5 0 1 1-.998-.06l.5-8.5a.5.5 0 0 1 .528-.47M8 4.5a.5.5 0 0 1 .5.5v8.5a.5.5 0 0 1-1 0V5a.5.5 0 0 1 .5-.5"/>
                        </svg>
                      </Button>
                  </td>
              </tr> 
            </tbody>

          </table>
        </div>
      </div>
    </div>

</template>  

<script>
import axios from 'axios';

export default {
  data() {
    return {
      items: [] 
    };
  },
  methods: {
    fetchData() {
      axios.get('http://127.0.0.1:8000/api/murid')
        .then(response => {
          this.items = response.data; 
        })
        .catch(error => {
          console.error('Error fetching data:', error);
        });
    },
    deleteItem(id) {
      if (confirm('Are you sure you want to delete this item?')) {
        axios.delete(`http://127.0.0.1:8000/api/murid/${id}`)
          .then(response => {
            console.log('Item deleted successfully:', response.data);
            this.fetchData();

            
          })
          .catch(error => {
            console.error('Error deleting item:', error);
          });
      }
    }
  },
  mounted() {
    this.fetchData();
  }
};



</script>
