<template>
  <div>
      <div class="flex">
        <div class="text-blue-600 border p-3 rounded border-blue-600 transition hover:bg-blue-600 hover:text-white cursor-pointer mr-5" @click="setImg" v-if="admin">Редактировать</div>
        <div class="text-red-600 border p-3 rounded border-red-600 transition hover:bg-red-600 hover:text-white cursor-pointer mr-5" @click="showDelete = true" v-if="admin && !showDelete">Удалить</div>
        <div class="text-red-600 p-3 mr-5" v-if="admin && showDelete">Вы уверены?</div>
        <div class="text-red-600 border p-3 rounded border-red-600 transition hover:bg-red-600 hover:text-white cursor-pointer mr-5" @click="deleteFromDb('deleteUnit', $route.params.id)" v-if="admin && showDelete">Да</div>
      </div>
      <div v-for="doc in jsonData" :key="doc.id" class="flex flex-col">
        <form v-if="edit" class="flex flex-col" @submit.prevent="updateToDb('updateUnit', doc.title, doc.name, doc.body, fileUrl, id)">
            <label for="image">Выберите картинку</label>
            <input @change="handleChange" class="mb-5" name="image" type="file" placeholder="Картинка">
            <input class="p-2 rounded shadow border-2 border-blue-600 ring-offset-2 mb-5" v-model="doc.title" type="text" placeholder="Заголовок">
            <input class="p-2 rounded shadow border-2 border-blue-600 ring-offset-2 mb-5" v-model="doc.name" type="text" placeholder="Тег">
            <textarea class="p-2 rounded shadow border-2 border-blue-600 ring-offset-2 mb-5" v-model="doc.body" type="text" placeholder="Текст"></textarea>

            <div>
                <button class="text-left p-3 rounded text-blue-600 transition bg-white border-2 border-blue-600 hover:bg-blue-600 hover:text-white" v-if="!isLoading">Сохранить</button>
                <button class="text-left p-3 rounded text-blue-600 transition bg-white border-2 border-blue-600 hover:bg-blue-600 hover:text-white" v-else disabled>Загрузка...</button>        
            </div>
        </form>
        </div>
      <div class="mt-5 grid lg:grid-cols-4 gap-x-10">
        <div class="col-span-4 md:col-span-3 mt-2">
          <div v-for="doc in jsonData" :key="doc.id">
              <div class="col-span-4 md:col-span-3 mt-2">
                <div class="font-nova-bold text-4xl">{{ doc.title }}</div>
                <div class="mt-10 text-lg" v-html="doc.body"></div>
              </div>
          </div>
          <div class="mt-10">
            <div class="flex">
              <div v-if="admin" @click="showForm = !showForm" class="p-2 mb-5 text-blue-600 border-2 border-blue-600 transition cursor-pointer rounded-md hover:text-white hover:bg-blue-600">Добавить направление исследований</div>
            </div>
            <form v-if="showForm" class="flex flex-col border-b border-blue-600 pb-10" @submit.prevent="writeToDb('createResearchDirection', researchText, id)">
              <textarea class="p-2 rounded shadow border-2 border-blue-600 ring-offset-2 mb-5" v-model="researchText" type="text" placeholder="Текст"></textarea>
              <div>
                  <button class="text-left p-3 rounded text-blue-600 transition bg-white border-2 border-blue-600 hover:bg-blue-600 hover:text-white" v-if="!isLoading">Сохранить</button>
                  <button class="text-left p-3 rounded text-blue-600 transition bg-white border-2 border-blue-600 hover:bg-blue-600 hover:text-white" v-else disabled>Загрузка...</button>        
              </div>
            </form>

            <h1 v-if="records.length > 0" class="text-2xl font-nova-bold mb-5">Направления исследований</h1>
            <div v-for="record in records[0]" :key="record.id" class="flex items-center mb-5" @click.alt="deleteFromDb('deleteResearchDirection', record.id.toString())">
              <div class="bg-blue-600 p-2 rounded-full mr-3"></div>
              <p class="text-lg" v-html="record.body"></p>
            </div>
          </div>

          <div class="col-span-3 hidden lg:flex mt-10">
              <table class="table-fixed w-full max-h-16">
                <thead class="shadow-brand sticky top-3 rounded-lg">
                    <tr class=" overscroll-x-auto">
                        <th class="w-1/4 text-left pl-5 py-5">Должность</th>
                        <th class="w-1/4 text-left pl-5 py-5">ФИО</th>
                        <th class="w-1/4 text-left pl-5 py-5">Телефон</th>
                        <th class="w-1/4 text-left pl-5 py-5">E-mail</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="person in jsonData" :key="person.id" @click.alt="deleteFromDb('deleteManagement', person.id)" class="stag relative h-20 border-b border-black">
                        <td class="pl-5 py-3 select-all">{{ person.position }}</td>
                        <td class="pl-5 py-3 select-all text-blue-600">
                          <router-link :to="`/management/${person.id}`">{{ person.name }}</router-link>
                        </td>
                        <td class="pl-5 py-3 select-all">{{ person.phone }}</td>
                        <td class="pl-5 py-3 select-all">{{ person.email }}</td>
                    </tr>
                </tbody>
               </table>
          </div>

          <div class="col-span-4 sm:grid-cols-2 lg:col-span-3 grid gap-5 lg:hidden">
              <div v-for="person in jsonData" :key="person.id" class="shadow-md p-3 rounded-lg">
                  <div class="flex justify-between">
                      <div class="font-nova-semi">Должность</div>
                      <div class="text-sm text-right">{{ person.position }}</div>
                  </div>
                  <div class="flex justify-between mt-2">
                      <div class="font-nova-semi">ФИО</div>
                      <div class="text-sm text-right">
                        <router-link :to="`/management/${person.id}`">{{ person.name }}</router-link>
                      </div>
                  </div>
                  <div class="flex justify-between mt-2">
                      <div class="font-nova-semi">Телефон</div>
                      <div class="text-sm text-right">{{ person.phone }}</div>
                  </div>
                  <div class="flex justify-between mt-2">
                      <div class="font-nova-semi">E-mail</div>
                      <div class="text-sm text-right">{{ person.email }}</div>
                  </div>
              </div>
          </div>
        </div>
        <div class="mt-10 lg:mt-0 col-span-4 lg:col-span-1 mb-5">
            <div class="grid sm:grid-cols-2 gap-5 lg:grid-cols-1 sticky top-3">
                <SingleNews class="mb-5"></SingleNews>
                <SingleAd class="mb-5"></SingleAd>
            </div>
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

const route = useRoute()
const { deleteFromDb } = deleteRecord()
const { writeToDb } = createRecord()
const { admin, checkAuth } = getAuth()
const { getJson, jsonData } = readRecords() 
const { getSubRecords } = readSubRecords()
const { updateToDb, handleChange, fileUrl } = updateRecord()
const id = route.params.id
console.log(id)
checkAuth()

const setImg = () => {
    edit.value = !edit.value
    let i = document.getElementById("imgPath")
    console.log(i.getAttribute("src"))
    fileUrl.value = i.getAttribute("src")
}

getJson('singleUnit', id)
let records = getSubRecords('researchDirections', id)

const showForm = ref(false)
const researchText = ref('')

const showDelete = ref(false)
const edit = ref(false)
const singleNews = ref('')
</script>
