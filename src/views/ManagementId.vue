<template>
  <div>
      <div class="flex">
        <div class="text-blue-600 border p-3 rounded border-blue-600 transition hover:bg-blue-600 hover:text-white cursor-pointer mr-5" @click="setImg" v-if="isAdmin || username == jsonData[0].created_by">Редактировать</div>
        <div class="text-red-600 border p-3 rounded border-red-600 transition hover:bg-red-600 hover:text-white cursor-pointer mr-5" @click="showDelete = true" v-if="(isAdmin || username == jsonData[0].created_by) && !showDelete">Удалить</div>
        <div class="text-red-600 p-3 mr-5" v-if="username && showDelete">Вы уверены?</div>
        <div class="text-red-600 border p-3 rounded border-red-600 transition hover:bg-red-600 hover:text-white cursor-pointer mr-5" @click="deleteFromDb('deleteManagement', $route.params.id)" v-if="username && showDelete">Да</div>
      </div>
      <div v-for="doc in jsonData" :key="doc.id" class="flex flex-col">
        <form v-if="edit" class="flex flex-col" @submit.prevent="updateToDb('updateManagement', doc.name, doc.position, doc.phone, doc.email, doc.unit, doc.about, doc.links, fileUrl, id)">
            <label for="image">Выберите картинку</label>
            <input @change="handleChange" class="mb-5" name="image" type="file" placeholder="Картинка">
            <input class="p-2 rounded shadow border-2 border-blue-600 ring-offset-2 mb-5" v-model="doc.name" type="text" placeholder="Имя">
            <input class="p-2 rounded shadow border-2 border-blue-600 ring-offset-2 mb-5" v-model="doc.position" type="text" placeholder="Должность">
            <input class="p-2 rounded shadow border-2 border-blue-600 ring-offset-2 mb-5" v-model="doc.phone" type="text" placeholder="Телефон">
            <input class="p-2 rounded shadow border-2 border-blue-600 ring-offset-2 mb-5" v-model="doc.email" type="text" placeholder="E-mail">
            <input class="p-2 rounded shadow border-2 border-blue-600 ring-offset-2 mb-5" v-model="doc.unit" type="text" placeholder="Подразделение">
            <textarea class="p-2 rounded shadow border-2 border-blue-600 ring-offset-2 mb-5" v-model="doc.about" type="text" placeholder="Описание"></textarea>
            <input class="p-2 rounded shadow border-2 border-blue-600 ring-offset-2 mb-5" v-model="doc.links" type="text" placeholder="Соцсети">

            <div>
                <button class="text-left p-3 rounded text-blue-600 transition bg-white border-2 border-blue-600 hover:bg-blue-600 hover:text-white" v-if="!isLoading">Сохранить</button>
                <button class="text-left p-3 rounded text-blue-600 transition bg-white border-2 border-blue-600 hover:bg-blue-600 hover:text-white" v-else disabled>Загрузка...</button>        
            </div>
        </form>
      </div>
      <div class="grid lg:grid-cols-4 gap-x-10">
        <div class="col-span-4 mt-2">
          <div class="mb-10 font-nova-bold text-3xl">Сотрудники ИДВ РАН</div>
          <div v-for="doc in jsonData" :key="doc.id">
              <div class="flex space-x-5">
                <img v-if="doc.img" class="object-cover w-28 h-28 rounded-full" :src="doc.img" alt="">
                <div>
                  <div class="text-blue-600 text-xl font-nova-semi">{{ doc.name }}</div>
                  <div class="mb-2">{{ doc.position }}</div>
                  <div>{{ doc.phone }}</div>
                  <div>{{ doc.email }}</div>
                </div>
              </div>
              <div class="text-lg mt-5" v-html="doc.about"></div>
          </div>

          <tabs v-model="active">
            <tab title="Основное" class="space-y-5">
              <Education :id="id"/>
              <EducationPlus :id="id"/>
              <Achievements :id="id" />
            </tab>
            <tab title="Преподавание">Преподавание</tab>
            <tab title="Публикации и исследование">Публикации и исследование</tab>
            <tab title="Опыт работы">Опыт работы</tab>
            <tab title="Прочее">Прочее</tab>
          </tabs>


          

          <!-- <div class="flex">
            <div v-if="admin" @click="showForm = !showForm" class="p-2 mt-5 text-blue-600 border-2 border-blue-600 transition cursor-pointer rounded-md hover:text-white hover:bg-blue-600">Добавить статью</div>
          </div>
          <form v-if="showForm" class="flex flex-col border-b border-blue-600 pb-10" @submit.prevent="writeToDb('createPaper', paper.title, paper.url, id)">
            <input class="p-2 rounded shadow border-2 border-blue-600 ring-offset-2 mb-5" v-model="paper.title" type="text" placeholder="Название">
            <input class="p-2 rounded shadow border-2 border-blue-600 ring-offset-2 mb-5" v-model="paper.url" type="text" placeholder="Ссылка">
            <div>
                <button class="text-left p-3 rounded text-blue-600 transition bg-white border-2 border-blue-600 hover:bg-blue-600 hover:text-white" v-if="!isLoading">Сохранить</button>
                <button class="text-left p-3 rounded text-blue-600 transition bg-white border-2 border-blue-600 hover:bg-blue-600 hover:text-white" v-else disabled>Загрузка...</button>        
            </div>
          </form>
          <div v-if="papers.length > 0" class="mt-10 mb-5 font-nova-bold text-2xl">Дополнительное образование / Повышение квалификации / Стажировки</div>
          <div class="grid grid-cols-1">
            <div class="mb-5" v-for="paper in papers[0]" :key="paper.id" @click.alt="deleteFromDb('deletePaper', paper.id.toString())">
              <router-link :to="paper.paper_url" class="hover:text-blue-600">
                <div>{{ paper.paper_title }}</div>
              </router-link>
            </div>
          </div> -->

          <!-- <div class="flex">
            <div v-if="admin" @click="showForm = !showForm" class="p-2 mt-5 text-blue-600 border-2 border-blue-600 transition cursor-pointer rounded-md hover:text-white hover:bg-blue-600">Добавить достижение</div>
          </div>
          <form v-if="showForm" class="flex flex-col border-b border-blue-600 pb-10" @submit.prevent="writeToDb('createPaper', paper.title, paper.url, id)">
            <input class="p-2 rounded shadow border-2 border-blue-600 ring-offset-2 mb-5" v-model="paper.title" type="text" placeholder="Название">
            <input class="p-2 rounded shadow border-2 border-blue-600 ring-offset-2 mb-5" v-model="paper.url" type="text" placeholder="Ссылка">
            <div>
                <button class="text-left p-3 rounded text-blue-600 transition bg-white border-2 border-blue-600 hover:bg-blue-600 hover:text-white" v-if="!isLoading">Сохранить</button>
                <button class="text-left p-3 rounded text-blue-600 transition bg-white border-2 border-blue-600 hover:bg-blue-600 hover:text-white" v-else disabled>Загрузка...</button>        
            </div>
          </form>
          <div v-if="papers.length > 0" class="mt-10 mb-5 font-nova-bold text-2xl">Достижения и поощрения</div>
           -->
        </div>
      </div>
  </div>
</template>

<script setup>
import axios from 'axios'
import SingleNews from '../components/SingleNews.vue'
import SingleAd from '../components/SingleAd.vue'
import { onMounted, ref } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import deleteRecord from '../composables/deleteRecord'
import getAuth from '../composables/getAuth'
import readRecords from '../composables/readRecords'
import updateRecord from '../composables/updateRecord'
import createRecord from '../composables/createRecord'
import readSubRecords from '../composables/readSubRecords'
import Tab from '../components/Tab.vue'
import Tabs from '../components/Tabs.vue'
import Education from '../components/person/Education.vue'
import EducationPlus from '../components/person/EducationPlus.vue'
import Achievements from '../components/person/Achievements.vue'

const active = ref(1);

const route = useRoute()
const { deleteFromDb } = deleteRecord()
const { writeToDb } = createRecord()
const { isAdmin, username, uid, checkAuth } = getAuth()
const { getJson, jsonData } = readRecords() 
const { getSubRecords } = readSubRecords()
const { updateToDb, handleChange, fileUrl } = updateRecord()
const id = route.params.id
console.log(id)
checkAuth()

const setImg = () => {
    edit.value = !edit.value
    fileUrl.value = jsonData.value[0].img
}

const subRecords = []

const doc = ref({
  title: '',
})
const paper = ref({
  title: '',
  url: '',
})

getJson('singleManagement', id)

const showForm = ref(false)
const researchText = ref('')

const showDelete = ref(false)
const edit = ref(false)
const singleNews = ref('')
</script>
