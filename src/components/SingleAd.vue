<template>
  <div>
      <div class="flex justify-between mb-3">
          <div class="uppercase text-lg font-nova-semi">анонсы</div>
          <router-link to="/ads" class="text-blue-600">все анонсы</router-link>
      </div>
      <div v-for="singleAd in jsonData" :key="singleAd.id">
        <router-link  :to="`/ads/${singleAd.id}`">
            <div class="font-nova-bold mt-1" v-html="singleAd.time"></div>
            <div class="text-blue-600 text-xs font-nova-bold mt-2" v-html="singleAd.tag"></div>
            <div class="font-nova-bold mt-1 text-sm" v-html="singleAd.title"></div>
            <div class="mt-1 text-sm" v-html="singleAd.date"></div>
        </router-link>
      </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import readRecords from '../composables/readRecords'

const { getJson, jsonData } = readRecords()
getJson('singleAd', '1')

</script>

<style>

</style>