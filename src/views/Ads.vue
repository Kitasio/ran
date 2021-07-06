<template>
  <div>
      <div class="flex justify-between">
        <h1 class="text-3xl font-nova-bold">Анонсы и события института</h1>
        <div v-if="username" @click="showForm = !showForm" class="p-2 text-blue-600 border-2 border-blue-600 transition cursor-pointer rounded-md hover:text-white hover:bg-blue-600">Добавить анонс</div>
      </div>

      <form v-if="showForm" class="flex flex-col border-b border-blue-600 pb-10" @submit.prevent="writeToDb('createAd', doc.title, doc.tag, doc.date, doc.time, username)">
        
        <input class="p-2 rounded shadow border-2 border-blue-600 ring-offset-2 mb-5" v-model="doc.title" type="text" placeholder="Заголовок">
        <input class="p-2 rounded shadow border-2 border-blue-600 ring-offset-2 mb-5" v-model="doc.tag" type="text" placeholder="Тег">
        <input class="p-2 rounded shadow border-2 border-blue-600 ring-offset-2 mb-5" v-model="doc.date" type="text" placeholder="Дата">
        <input class="p-2 rounded shadow border-2 border-blue-600 ring-offset-2 mb-5" v-model="doc.time" type="text" placeholder="Время">

        <div>
            <button class="text-left p-3 rounded text-blue-600 transition bg-white border-2 border-blue-600 hover:bg-blue-600 hover:text-white" v-if="!isLoading">Сохранить</button>
            <button class="text-left p-3 rounded text-blue-600 transition bg-white border-2 border-blue-600 hover:bg-blue-600 hover:text-white" v-else disabled>Загрузка...</button>        
        </div>
      </form>

      <div class="mt-10 grid lg:grid-rows-2 gap-10 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4">
        <div class="relative" v-for="singleAd in jsonData" :key="singleAd.id">
          <router-link :to="`/ads/${singleAd.id}`">
            <div class="text-sm font-nova-bold mt-2" v-html="singleAd.time"></div>
            <div class="text-blue-600 text-xs font-nova-bold" v-html="singleAd.tag"></div>
            <div class="font-nova-bold mt-1 leading-none" v-html="singleAd.title"></div>
              <div class="text-xs" v-html="singleAd.date"></div>
          </router-link>
        </div>
      </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import createRecord from '../composables/createRecord'
import getAuth from '../composables/getAuth'
import readRecords from '../composables/readRecords'

const { getJson, jsonData } = readRecords()
getJson('allAds')
const showForm = ref(false)

const { handleChange, writeToDb, uploadImage, fileUrl } = createRecord()
const { isAdmin, username, uid, checkAuth } = getAuth()
checkAuth()

const doc = ref({
    title: '',
    tag: '',
    date: '',
    time: '',
})

</script>

<style>

</style>