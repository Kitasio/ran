import { ref } from 'vue'
import axios from 'axios'
import clientPath  from './config'

const readSubRecords = () => {
    const subRecords = ref('')
    const getSubRecords = (query, ...data) => {
        axios
            .post(clientPath + "/api/readSubRecords", {
                query: query,
                data: data,
            })
            .then(response => {
                subRecords.value = JSON.parse(response.data)
            })
            .catch(err => {
                console.log(err)
            })
    }

    return { getSubRecords, subRecords }
}

export default readSubRecords
