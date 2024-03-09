<script setup>
import Navbar from './../layout/mainView.vue'
</script>

<template>
  <Navbar/>

    <div class="mt-20 sm:mx-2 grid grid-cols-1 place-items-center">
      <div class="">

        <div class="relative overflow-x-auto shadow-md sm:rounded-lg">
          <table class="w-full text-sm text-left rtl:text-right text-gray-500 dark:text-gray-400">
            <thead class="text-xs text-gray-700 bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
              <tr>
                <th class="px-6 py-3"></th>
                <th class="px-6 py-3">Email</th>
                <th class="px-6 py-3">Nama</th>
                <th class="px-6 py-3">Kelas</th>
                <th class="px-6 py-3">NIS</th>
                <th class="px-6 py-3">Action</th>
              </tr>
            </thead>
            <tbody>
              <tr class="bg-white border-b dark:bg-gray-800 dark:border-gray-700">
                  <td class="px-6 py-4">1</td>
                  <td class="px-6 py-4 text-center">muhammad.ilyaasa@gmail.com</td>
                  <th class="px-6 py-4 font-medium text-gray-900 whitespace-nowrap dark:text-white">
                    Muhammad Ilyasa
                  </th>
                  <td class="px-6 py-4 text-center">XII RPL</td>
                  <td class="px-6 py-4 text-center">118695552</td>
                  <td class="px-6 py-4 text-center">
                      <a href="#" class="font-medium text-blue-600 dark:text-blue-500 hover:underline">Edit</a>
                  </td>
              </tr> 
            </tbody>

          </table>
        </div>
      </div>
    </div>

</template>  
