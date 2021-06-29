<template>
  <div>
      <div class="flex justify-between">
        <h1 class="text-3xl font-nova-bold">Дирекция института</h1>
        <div v-if="admin" @click="showForm = !showForm" class="p-2 text-blue-600 border-2 border-blue-600 transition cursor-pointer rounded-md hover:text-white hover:bg-blue-600">Добавить дирекцию</div>
      </div>

      <form v-if="showForm" class="flex flex-col border-b border-blue-600 pb-10" @submit.prevent="writeToDb('createManagement', doc.name, doc.position, doc.phone, doc.email, doc.unit)">
        <!-- <label for="image">Выберите картинку</label> -->
        <!-- <input @change="handleChange" class="mb-5" name="image" type="file" placeholder="Картинка"> -->
        <input class="p-2 rounded shadow border-2 border-blue-600 ring-offset-2 mb-5" v-model="doc.name" type="text" placeholder="Имя">
        <input class="p-2 rounded shadow border-2 border-blue-600 ring-offset-2 mb-5" v-model="doc.position" type="text" placeholder="Должность">
        <input class="p-2 rounded shadow border-2 border-blue-600 ring-offset-2 mb-5" v-model="doc.phone" type="text" placeholder="Телефон">
        <input class="p-2 rounded shadow border-2 border-blue-600 ring-offset-2 mb-5" v-model="doc.email" type="text" placeholder="E-mail">
        <input class="p-2 rounded shadow border-2 border-blue-600 ring-offset-2 mb-5" v-model="doc.unit" type="text" placeholder="Подразделение">

        <div>
            <button class="text-left p-3 rounded text-blue-600 transition bg-white border-2 border-blue-600 hover:bg-blue-600 hover:text-white">Сохранить</button>
        </div>
      </form>

      <div class="mt-5 grid lg:grid-cols-4 gap-x-10">
          <div class="col-span-3 hidden lg:flex">
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
          <div class="mt-10 lg:mt-0 col-span-4 lg:col-span-1 mb-5">
              <div class="grid sm:grid-cols-2 gap-5 lg:grid-cols-1 sticky top-3">
                  <SingleNews class="mb-5"></SingleNews>
                  <SingleAd></SingleAd>
              </div>
          </div>
      </div>
  </div>
</template>

<script setup>
import { ref } from "vue"
import SingleNews from '../components/SingleNews.vue'
import SingleAd from '../components/SingleAd.vue'
import readRecords from "../composables/readRecords"
import createRecord from "../composables/createRecord"
import deleteRecord from "../composables/deleteRecord"
import getAuth from "../composables/getAuth"

const doc = ref({
    name: '',
    position: '',
    phone: '',
    email: '',
    unit: '',
})

const del = (id) => {
    console.log(id)
}

const showForm = ref(false)
const { getJson, jsonData } = readRecords()
const { writeToDb } = createRecord()
const { deleteFromDb } = deleteRecord()
const { admin, checkAuth } = getAuth()

checkAuth()
getJson('getManagement')

</script>

<style>

</style>