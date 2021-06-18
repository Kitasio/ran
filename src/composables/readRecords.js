import { ref } from 'vue'
import axios from 'axios'

const readRecords = () => {
    const jsonData = ref('')
    const getJson = (query, ...data) => {
        axios
            .post("http://localhost:3090/api/readRecords", {
                query: query,
                data: data,
            })
            .then(response => {
                jsonData.value = JSON.parse(response.data)
            })
    }

    return { getJson, jsonData }
}

export default readRecords
