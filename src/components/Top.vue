<template>
  <div class="bg-blue-700 font-nova-semi text-sm text-white flex items-center  py-3">
      <h1>ИНСТИТУТ ДАЛЬНЕГО ВОСТОКА РАН | ОСНОВАН В 1966 ГОДУ</h1>
      <div class="ml-10" v-if="username">{{ username }}</div>
      <div class="ml-10 cursor-pointer" @click="logout()" v-if="username">Logout</div>
  </div>
</template>

<script setup>
import getAuth from "../composables/getAuth"
import signOut from "../composables/signOut"

const { isAdmin, username, uid, checkAuth } = getAuth()
const { logout } = signOut()
checkAuth()
</script>

<style>

</style>