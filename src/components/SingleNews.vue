<template>
  <div>
      <div class="flex justify-between mb-3">
          <div class="uppercase text-lg font-nova-semi">новости</div>
          <router-link to="/news" class="text-blue-600">все новости</router-link>
      </div>
      <div v-for="singleNews in jsonData" :key="singleNews.id">
        <router-link  :to="`/news/${singleNews.id}`">
            <img v-if="singleNews.img" class="w-full rounded" :src="singleNews.img" alt="">
            <div class="text-blue-600 text-xs font-nova-bold mt-2" v-html="singleNews.tag"></div>
            <div class="font-nova-bold mt-1 text-sm" v-html="singleNews.title"></div>
            <div class="mt-1 text-sm" v-html="singleNews.date"></div>
        </router-link>
      </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import readRecords from '../composables/readRecords'

const { getJson, jsonData } = readRecords()
getJson('singleNews', '11')

</script>

<style>

</style>