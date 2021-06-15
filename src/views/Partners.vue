<template>
  <div>
      <div class="flex justify-between">
        <h1 class="text-3xl font-nova-bold">Партнеры</h1>
        <div v-if="admin" @click="showForm = !showForm" class="p-2 text-blue-600 border-2 border-blue-600 transition cursor-pointer rounded-md hover:text-white hover:bg-blue-600">Добавить партнера</div>
      </div>

      <form v-if="showForm" class="flex flex-col border-b border-blue-600 pb-10" @submit.prevent="writeToDb('createPartner', doc.name, doc.link, fileUrl)">
        <label for="image">Выберите картинку</label>
        <div class="flex m-4">
            <input @change="handleChange($event)" class="mb-5 mr-5" name="file" type="file" placeholder="Картинка">
            <div v-if="!fileUrl" @click="uploadImage" class="p-2 text-blue-600 border-2 border-blue-600 transition cursor-pointer rounded-md hover:text-white hover:bg-blue-600">Сохранить</div>
            <div v-if="fileUrl" @click="uploadImage" class="p-2 transition cursor-pointer rounded-md text-blue-600">Готово</div>
        </div>
        <input class="p-2 rounded shadow border-2 border-blue-600 ring-offset-2 mb-5" v-model="doc.name" type="text" placeholder="Имя">
        <input class="p-2 rounded shadow border-2 border-blue-600 ring-offset-2 mb-5" v-model="doc.link" type="text" placeholder="Ссылка">

        <div>
            <button class="text-left p-3 rounded text-blue-600 transition bg-white border-2 border-blue-600 hover:bg-blue-600 hover:text-white" v-if="!isLoading">Сохранить</button>
            <button class="text-left p-3 rounded text-blue-600 transition bg-white border-2 border-blue-600 hover:bg-blue-600 hover:text-white" v-else disabled>Загрузка...</button>        
        </div>
      </form>

      <div class="mt-5 grid lg:grid-cols-4 gap-x-14 text-lg">
          <div class="col-span-4 lg:col-span-3">
              <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
                  <div class="relative w-full" v-for="partner in jsonData" :key="partner.id" @click.alt="deleteFromDb('deletePartner', partner.id)">
                    <a :href="partner.link" class="flex h-full flex-col justify-between rounded-lg transition duration-500 shadow-brand hover:shadow-bigger p-3">
                        <div class="flex h-32 items-center justify-center">
                            <img class="max-h-28" :src="partner.img" alt="">
                        </div>
                        <div class="text-blue-600" v-html="partner.name"></div>
                    </a>
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
import deleteRecord from '../composables/deleteRecord'

const { getJson, jsonData } = readRecords()
getJson('allPartners')
const showForm = ref(false)
const { deleteFromDb } = deleteRecord()

const { handleChange, writeToDb, uploadImage, fileUrl } = createRecord()
const { admin, checkAuth } = getAuth()
checkAuth()

const doc = ref({
    name: '',
    img: '',
    link: '',
})
</script>

<style>

</style>