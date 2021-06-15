<template>
  <div class="flex flex-col justify-between">
      <div class="flex flex-col">
          <div class="flex justify-center mt-5">
              <img src="/images/sidebar/ru.svg" alt="">
              <img class="mx-3" src="../assets/images/sidebar/ch.svg" alt="">
              <img src="../assets/images/sidebar/uk.svg" alt="">
          </div>
          <div class="flex justify-center mt-5">
              <img src="../assets/images/sidebar/logo.png" alt="">
          </div>
          <div class="mt-5 ml-8 mr-3">
              <div v-if="admin" class="my-3">
                <router-link to="/admin" class="text-md uppercase font-semibold hover:text-blue-600">Админ</router-link>
              </div>
              <div class="my-3">
                <router-link to="/" class="text-md uppercase font-semibold hover:text-blue-600">Главная</router-link>
              </div>
              <div>
                <div class="my-3 flex justify-between">
                    <h1 to="/" class="text-md uppercase font-semibold hover:text-blue-600 cursor-pointer" @click="links1.show = !links1.show">Институт</h1>
                    <img @click="links1.show = !links1.show" class="w-5 mr-7 cursor-pointer" src="../assets/images/icons/chevron-down.svg" alt="">
                </div>
                <div v-if="links1.show" class="ml-3 flex flex-col">
                    <router-link :to="link.link" class="mb-1 hover:text-blue-600" v-for="(link, index) in links1.items" :key="index">
                        {{ link.name }}
                    </router-link>
                </div>
              </div>

              <div>
                <div class="my-3 flex justify-between">
                    <h1 to="/" class="text-md uppercase font-semibold hover:text-blue-600 cursor-pointer" @click="links2.show = !links2.show">ОБРАЗОВАНИЕ</h1>
                    <img @click="links2.show = !links2.show" class="w-5 mr-7 cursor-pointer" src="../assets/images/icons/chevron-down.svg" alt="">
                </div>
                <div v-if="links2.show" class="ml-3 flex flex-col">
                    <router-link to="/" class="mb-1 hover:text-blue-600" v-for="(link, index) in links2.items" :key="index">
                        {{ link }}
                    </router-link>
                </div>
              </div>

              <div>
                <div class="my-3 flex justify-between">
                    <h1 to="/" class="text-md uppercase font-semibold hover:text-blue-600 cursor-pointer" @click="links3.show = !links3.show">НАУКА</h1>
                    <img @click="links3.show = !links3.show" class="w-5 mr-7 cursor-pointer" src="../assets/images/icons/chevron-down.svg" alt="">
                </div>
                <div v-if="links3.show" class="ml-3 flex flex-col">
                    <router-link to="/" class="mb-1 hover:text-blue-600" v-for="(link, index) in links3.items" :key="index">
                        {{ link }}
                    </router-link>
                </div>
              </div>
              <div class="my-3">
                <router-link to="/#########" class="text-md uppercase font-semibold hover:text-blue-600">Контакты</router-link>
              </div>
          </div>
      </div>

      <div>
        <div class="ml-8 bottom-20 font-thin text-sm text-gray-400 mx-10">Copyright © 2020<br>ИДВ РАН</div>
        <img class="" src="../assets/images/sidebar/bot-icon.svg" alt="">
      </div>
  </div>
</template>

<script>
export default {
    data() {
        return {
            admin: null,
            links1: {
                show: false,
                items: [
                    { name: 'О нас', link: '/about' },
                    { name: 'Дирекция', link: '/management' },
                    { name: 'Подразделения', link: '/units' },
                    { name: 'Ученый совет', link: '/council' },
                    { name: 'Документы', link: '/documents' },
                    { name: 'Партнеры', link: '/partners' },
                    { name: 'Вакансии и конкурсы', link: '/#####' },
                ]
            },
            links2: {
                show: false,
                items: [
                    'Аспирантура',
                    'Диссертационный совет',
                ]
            },
            links3: {
                show: false,
                items: [
                    'События',
                    'Издания',
                    'Библиотеки и базы данных',
                    'Проекты',
                ]
            },
        }
    }
}
</script>

<style>

</style>