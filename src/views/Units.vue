<template>
  <div>
      <div class="flex justify-between">
        <h1 class="text-3xl font-nova-bold">Подразделения</h1>
        <div v-if="admin" @click="showForm = !showForm" class="p-2 text-blue-600 border-2 border-blue-600 transition cursor-pointer rounded-md hover:text-white hover:bg-blue-600">Добавить подразделение</div>
      </div>

      <form v-if="showForm" class="flex flex-col border-b border-blue-600 pb-10" @submit.prevent="writeToDb('createUnit', doc.title, doc.name, doc.body, fileUrl)">
        <label for="image">Выберите картинку</label>
        <div class="flex m-4">
            <input @change="handleChange($event)" class="mb-5 mr-5" name="file" type="file" placeholder="Картинка">
            <div v-if="!fileUrl" @click="uploadImage" class="p-2 text-blue-600 border-2 border-blue-600 transition cursor-pointer rounded-md hover:text-white hover:bg-blue-600">Сохранить</div>
            <div v-if="fileUrl" @click="uploadImage" class="p-2 transition cursor-pointer rounded-md text-blue-600">Готово</div>
        </div>
        <input class="p-2 rounded shadow border-2 border-blue-600 ring-offset-2 mb-5" v-model="doc.title" type="text" placeholder="Заголовок">
        <input class="p-2 rounded shadow border-2 border-blue-600 ring-offset-2 mb-5" v-model="doc.name" type="text" placeholder="Имя">
        <textarea class="p-2 rounded shadow border-2 border-blue-600 ring-offset-2 mb-5" v-model="doc.body" type="text" placeholder="Текст"></textarea>

        <div>
            <button class="text-left p-3 rounded text-blue-600 transition bg-white border-2 border-blue-600 hover:bg-blue-600 hover:text-white" v-if="!isLoading">Сохранить</button>
            <button class="text-left p-3 rounded text-blue-600 transition bg-white border-2 border-blue-600 hover:bg-blue-600 hover:text-white" v-else disabled>Загрузка...</button>        
        </div>
      </form>

      <div class="mt-5 grid lg:grid-cols-4 gap-x-14 text-lg">
          <div class="col-span-4 lg:col-span-3">
              <div class="stag mb-5">В составе Института функционируют следующие научно-исследовательские центры:</div>
              <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
                  <div class="relative w-full" v-for="unit in jsonData" :key="unit.id">
                    <router-link :to="`/units/${unit.id}`" class="flex h-full flex-col justify-between rounded-lg transition duration-500 shadow-brand hover:shadow-bigger p-3">
                        <div class="flex justify-between">
                            <div class="font-nova-semi mb-10" v-html="unit.title"></div>
                            <img class="w-7 self-start" :src="unit.img" alt="">
                        </div>
                        <div class="text-sm text-gray-500">{{ unit.name }}</div>
                    </router-link>
                  </div>
              </div>
          </div>
          <div class="mt-10 lg:mt-0 col-span-4 lg:col-span-1 mb-5">
              <div class="grid sm:grid-cols-2 gap-5 lg:grid-cols-1">
                  <SingleNews class="mb-5"></SingleNews>
                  <SingleAd></SingleAd>
              </div>
          </div>
      </div>
  </div>
</template>

<script setup>
import SingleNews from '../components/SingleNews.vue'
import SingleAd from '../components/SingleAd.vue'
import { ref } from 'vue'
import createRecord from '../composables/createRecord'
import getAuth from '../composables/getAuth'
import readRecords from '../composables/readRecords'

const { getJson, jsonData } = readRecords()
getJson('allUnits')
const showForm = ref(false)

const { handleChange, writeToDb, uploadImage, fileUrl } = createRecord()
const { admin, checkAuth } = getAuth()
checkAuth()

const doc = ref({
    title: '',
    name: '',
    img: '',
    body: '',
})
</script>

<style>

</style>