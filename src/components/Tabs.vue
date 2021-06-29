<script>
import { provide, computed, ref } from "vue";

export default {
    name: "Tabs",
    props: {
        modelValue: {
            type: [String, Number],
        },
    },
    emits: ["update:modelValue"],
    setup(props, { slots, emit }) {
        const active = computed(() => props.modelValue);
        const tabs = ref([]);

        function selectTab(tab) {
            emit("update:modelValue", tab);
        }

        provide("tabsState", {
            active,
            tabs,
        });

        return {
            tabs,
            active,
            selectTab,
        };
    },
};
</script>

<template>
    <ul class="flex flex-col lg:flex-row lg:space-x-10">
        <li
            v-for="(tab, i) of tabs"
            :key="i"
            :class="
                active === i
                    ? 'border-b-4 border-blue-600 box-content text-blue-600'
                    : 'border-gray-500 text-gray-500'
            "
            class="flex items-center px-0 py-3 rounded-tl-md rounded-tr-md overflow-hidden cursor-pointer hover:text-blue-800"
            @click="selectTab(i)"
        >
            {{ tab.props.title }}
        </li>
    </ul>
    <div class="bg-gray-300 -m-0.5 h-0.5"></div>
    <div class="mt-6">
        <slot />
    </div>
</template>