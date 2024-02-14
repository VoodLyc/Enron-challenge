import axios from 'axios';

class EmailService {
    constructor() {
        this.http = axios.create({
            baseURL: 'http://localhost:8080'
        });
    }

    async searchEmails(term) {
        try {
            const response = await this.http.get(`/search?term=${term}`);
            return response.data;
        } catch (err) {
            console.error('Error fetching data', err);
            throw err;
        }
    }
}

export default new EmailService();