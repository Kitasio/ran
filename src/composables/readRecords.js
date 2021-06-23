import { ref } from 'vue'
import axios from 'axios'
import clientPath  from './config'

const readRecords = () => {
    const jsonData = ref(null)
    const getJson = async (query, ...data) => {
        await axios
            .post(clientPath + "/api/readRecords", {
                query: query,
                data: data,
            })
            .then(response => {
                jsonData.value = JSON.parse(response.data)
            })
            .catch(err => {
                console.log(err)
            })
    }

    return { getJson, jsonData }
}

export default readRecords
