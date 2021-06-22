import { ref } from 'vue'
import axios from 'axios'
import clientPath  from './config'

const readSubRecords = () => {
    const getSubRecords = (query, ...data) => {
        const subRecords = ref([])
        axios
            .post(clientPath + "/api/readSubRecords", {
                query: query,
                data: data,
            })
            .then(response => {
                subRecords.value.push(JSON.parse(response.data)) 
            })
            .catch(err => {
                console.log(err)
            })
        return subRecords.value
    }

    return { getSubRecords }
}

export default readSubRecords
