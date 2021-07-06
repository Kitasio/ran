<template>
  <div>
      <div class="flex justify-between">
        <h1 class="text-3xl font-nova-bold">Новости института</h1>
        <div v-if="username" @click="showForm = !showForm" class="p-2 text-blue-600 border-2 border-blue-600 transition cursor-pointer rounded-md hover:text-white hover:bg-blue-600">Добавить новость</div>
      </div>

      <form v-if="showForm" class="flex flex-col border-b border-blue-600 pb-10" @submit.prevent="writeToDb('createNews', doc.title, doc.tag, doc.date, doc.body, fileUrl, username)">
        <label for="image">Выберите картинку</label>
        <input @change="handleChange($event)" class="mb-5 mr-5" name="file" type="file" placeholder="Картинка">
        <input class="p-2 rounded shadow border-2 border-blue-600 ring-offset-2 mb-5" v-model="doc.title" type="text" placeholder="Заголовок">
        <input class="p-2 rounded shadow border-2 border-blue-600 ring-offset-2 mb-5" v-model="doc.tag" type="text" placeholder="Тег">
        <input class="p-2 rounded shadow border-2 border-blue-600 ring-offset-2 mb-5" v-model="doc.date" type="text" placeholder="Дата">
        <textarea class="p-2 rounded shadow border-2 border-blue-600 ring-offset-2 mb-5" v-model="doc.body" type="text" placeholder="Текст"></textarea>

        <div>
            <button class="text-left p-3 rounded text-blue-600 transition bg-white border-2 border-blue-600 hover:bg-blue-600 hover:text-white">Сохранить</button>
        </div>
      </form>

      <div class="mt-10 grid lg:grid-rows-2 gap-10 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4">
        <div class="relative" v-for="singleNews in jsonData" :key="singleNews.id">
          <router-link :to="`/news/${singleNews.id}`">
            <img class="w-full rounded" :src="singleNews.img" alt="">
            <div class="text-blue-600 text-xs font-nova-bold mt-2" v-html="singleNews.tag"></div>
            <div class="font-nova-bold mt-1 text-sm" v-html="singleNews.title"></div>
            <div class="mt-1 text-sm" v-html="singleNews.date"></div>
          </router-link>
        </div>
      </div>

      <div class="flex mt-5">
        <!-- <div @click="prev" class="font-nova-semi text-lg cursor-pointer py-1 px-3 transition text-blue-600 hover:text-white hover:bg-blue-600 rounded">Назад</div> -->
        <div @click="loadMore" class="font-nova-semi text-lg cursor-pointer text-blue-600">Показать еще</div>
      </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import createRecord from '../composables/createRecord'
import getAuth from '../composables/getAuth'
import readRecords from '../composables/readRecords'

const { getJson, jsonData } = readRecords()
getJson('allNews')
const showForm = ref(false)

const { handleChange, writeToDb, uploadImage, fileUrl } = createRecord()
const { isAdmin, username, uid, checkAuth } = getAuth()
checkAuth()

const doc = ref({
    title: '',
    tag: '',
    date: '',
    body: '',
})

</script>

<style>

</style>