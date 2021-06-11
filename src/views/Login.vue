<template>
  <div>
      <form class="flex flex-col md:w-1/2" @submit.prevent="login()">
          <input class="p-2 mb-5 focus:bg-gray-100 rounded-md border-2 border-blue-600" v-model="username" type="text" required placeholder="username">
          <input class="p-2 mb-5 focus:bg-gray-100 rounded-md border-2 border-blue-600" v-model="password" type="password" required placeholder="password">
          <div>
            <button class="text-left p-2 border-2 border-blue-600 transition text-blue-600 hover:text-white hover:bg-blue-600 rounded-md">login</button>
          </div>
          <div class="mt-3 text-red-400">{{ error }}</div>
      </form>
      <div @click="getJson('getAllUsers')">get json data</div>
      <div v-for="item in jsonData" :key="item.uid">
        <div>{{ item.name }}</div>
        <div>{{ item.password }}</div>
        <div>{{ item.uid }}</div>
      </div>
  </div>
</template>

<script setup>
import getAuth from "../composables/getAuth"
import signIn from "../composables/signIn"
import signOut from "../composables/signOut"
import readRecords from "../composables/readRecords"

const { user, checkAuth } = getAuth()
const { username, password, error, login } = signIn()
const { logout } = signOut()
const { getJson, jsonData } = readRecords()

</script>

<style>

</style>