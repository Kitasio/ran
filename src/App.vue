<template>
  <div class="grid grid-cols-12 font-nova">
    <img @click="showSidebar = !showSidebar" class="fixed z-30 bottom-4 right-4 m-3 md:hidden w-10 cursor-pointer" src="./assets/images/icons/menu.svg" alt="">
    <div :class="{ 'fixed': showSidebar, 'hidden': !showSidebar }" class="bg-white z-30 md:static md:grid md:col-span-3 lg:col-span-2 min-h-screen shadow-big">
      <div class="">
        <Sidebar id="sidebar" class="sticky top-0 h-screen overflow-y-scroll"></Sidebar>
      </div>
    </div>
    <div class="col-span-12 md:col-span-9 lg:col-span-10">
      <div>
        <Top class="pl-5 lg:px-14 xl:px-20"></Top>
        <div class="px-5 lg:px-14 xl:px-20 py-10">
          <router-view></router-view>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import HelloWorld from './components/HelloWorld.vue'
import Top from './components/Top.vue'
import Sidebar from './components/Sidebar.vue'
import { ref } from 'vue'

const showSidebar = ref(false)
// This starter template is using Vue 3 experimental <script setup> SFCs
// Check out https://github.com/vuejs/rfcs/blob/script-setup-2/active-rfcs/0000-script-setup.md
</script>

<style>

</style>
