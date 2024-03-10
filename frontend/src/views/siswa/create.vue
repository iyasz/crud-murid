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
          <label for="email" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Email</label>
          <input v-model="formData.email" type="email" id="email" name="email" class="shadow-sm bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 focus:text-gray-900 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500 dark:shadow-sm-light" required />
        </div>
        <div class="mb-5">
          <label for="nis" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">NIS</label>
          <input v-model="formData.nis" type="number" id="nis" name="nis" class="shadow-sm bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 focus:text-gray-900 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500 dark:shadow-sm-light" required />
        </div>
        <div class="mb-5">
          <label for="class" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Kelas</label>
          <select v-model="formData.class_id" name="class_id" class="shadow-sm bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 focus:text-gray-900 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500 dark:shadow-sm-ligh" id="" required>
            <option selected disabled></option>

            <option v-for="data in datas" :key="data.id" :value="data.id">{{ data.name +' '+ data.jurusan }}</option>
          </select>
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
        email: '',
        nis: '',
        class_id: '',
      },
      datas: [],
    };
  },
  methods: {
    getDataClass(){
      axios.get("http://127.0.0.1:8000/api/class")
      .then(response => {
          this.datas = response.data; 
        })
        .catch(error => {
          console.error('Error fetching data:', error);
        });
    },
    submitForm() {
      axios.post('http://127.0.0.1:8000/api/murid', this.formData)
        .then(response => {
          console.log('Data inserted successfully:', response.data);
          // Reset the form after successful submission
          this.formData.name = '';
          this.formData.email = '';
          this.formData.nis = '';
          this.formData.class_id = '';

          this.$router.push({ name: 'murid.index' });
        })
        .catch(error => {
          console.error('Error inserting data:', error);
        });
    }
  },
  mounted() {
    this.getDataClass();
  }
};

</script>
