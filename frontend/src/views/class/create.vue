<script setup>
import Navbar from './../layout/mainView.vue'
</script>

<template>
  <Navbar/>

    <div class="mt-40 sm:mx-2 lg:mx-10">

      <form class="max-w-sm mx-auto" @submit.prevent="submitForm">
        <div class="mb-5">
          <label for="name" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Nama</label>
          <input v-model="formData.name" type="text" id="name" name="name" class="shadow-sm bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 focus:text-gray-900 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500 dark:shadow-sm-light" required />
        </div>
        <div class="mb-5">
          <label for="jurusan" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Jurusan</label>
          <input v-model="formData.jurusan" type="text" id="jurusan" name="jurusan" class="shadow-sm bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 focus:text-gray-900 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500 dark:shadow-sm-light" required />
        </div>

        <button type="submit" class="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">Submit</button>
      </form>

    </div>

</template>  

<script>
import axios from 'axios';

export default {
  data() {
    return {
      formData: {
        name: '',
        jurusan: ''
      }
    };
  },
  methods: {
    submitForm() {
      axios.post('http://127.0.0.1:8000/api/class', this.formData)
        .then(response => {
          console.log('Data inserted successfully:', response.data);
          // Reset the form after successful submission
          this.formData.name = '';
          this.formData.jurusan = '';

          this.$router.push({ name: 'classIndex' });
        })
        .catch(error => {
          console.error('Error inserting data:', error);
        });
    }
  }
};

</script>
