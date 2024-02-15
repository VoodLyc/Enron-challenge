import axios from 'axios';

class EmailService {
    constructor() {
        this.http = axios.create({
            baseURL: 'http://localhost:8080'
        });
    }

    async getEmails(from) {
        try {
            const response = await this.http.get(`/emails?from=${from}`);
            return response.data;
        } catch (err) {
            console.error('Error fetching data', err);
            throw err;
        } 
    }

    async searchEmails(term, from) {
        try {
            const response = await this.http.get(`/emails/search?term=${term}&from=${from}`);
            return response.data;
        } catch (err) {
            console.error('Error fetching data', err);
            throw err;
        }
    }
}

export default new EmailService();