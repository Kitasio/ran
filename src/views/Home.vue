<template>
  <div class="grid gap-10 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4">
      <router-link :to="`/news/${jsonData[0].id}`" class="flex flex-col md:flex-row col-span-full lg:col-span-3">
          <img class="md:w-1/2 rounded-md object-cover overflow-hidden border" :src="jsonData[0].img" alt="">
          <div class="md:ml-5">
              <div class="text-blue-600 text-lg font-nova-bold uppercase">{{ jsonData[0].tag }}</div>
              <div class="font-nova-bold mt-3 text-xl">{{ jsonData[0].title }}</div>
              <div class="text-sm text-gray-500">{{ jsonData[0].date }}</div>
          </div>
      </router-link>
      <div class="lg:col-span-1">
          <div class="flex justify-between items-center mb-5">
              <div class="font-nova-bold text-xl">Анонсы</div>
              <router-link to="/ads" class="text-blue-600">все анонсы</router-link>
          </div>
          <router-link :to="`/ads/${ad[0][0].id}`">
              <div class="text-lg font-nova-bold">{{ ad[0][0].time }}</div>
              <div class="text-blue-600 text-sm mb-1 font-nova-bold uppercase">{{ ad[0][0].tag }}</div>
              <div class="font-nova-bold text-sm mb-2">{{ ad[0][0].title }}</div>
              <div class="text-sm">{{ ad[0][0].date }}</div>
          </router-link>
      </div>
  </div>
    
  <div class="flex items-center mt-10 mb-3">
    <div class="font-nova-bold text-xl mr-10">Новости</div>
    <router-link to="/news" class="text-blue-600">все новости</router-link>
  </div>
  <div class="grid gap-10 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4">
    <div class="relative" v-for="singleNews in jsonData.slice(1,5)" :key="singleNews.id">
        <router-link :to="`/news/${singleNews.id}`">
            <img class="w-full rounded" :src="singleNews.img" alt="">
            <div class="text-blue-600 text-xs font-nova-bold mt-2" v-html="singleNews.tag"></div>
            <div class="font-nova-bold mt-1 text-sm" v-html="singleNews.title"></div>
            <div class="mt-1 text-sm text-gray-500" v-html="singleNews.date"></div>
        </router-link>
    </div>  
  </div>

  <div class="flex items-center mt-10 mb-3">
    <div class="font-nova-bold text-xl mr-10">Издания</div>
    <router-link to="/library" class="text-blue-600">все издания</router-link>
  </div>
  <div class="grid gap-10 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4">
      <div v-for="book in books[0]" :key="book.id" @click.alt="deleteFromDb('deleteBook', book.id.toString())">
        <div class="flex">
          <img class="w-20 object-cover" :src="book.book_img" alt="">
          <div class="ml-3 text-sm">{{ book.book_title }}</div>
        </div>
      </div>
      <div class="lg:col-span-1">
          <div class="flex flex-col justify-between h-full">
              <router-link to="/about">
                <img src="/src/assets/images/home/about.svg" alt="">
              </router-link>
              <router-link to="/graduate">
                <img src="/src/assets/images/home/graduate.svg" alt="">
              </router-link>
          </div>
      </div>
  </div>
</template>

<script setup>
import { onMounted } from "@vue/runtime-core";
import readRecords from "../composables/readRecords";
import readSubRecords from "../composables/readSubRecords";

const { getSubRecords } = readSubRecords()
const { getJson, jsonData } = readRecords()
getJson('homepageNews')
const ad = getSubRecords('latestAd')
const books = getSubRecords('getBooks')
</script>

<style>

</style>