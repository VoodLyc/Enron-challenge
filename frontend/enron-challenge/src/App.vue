<template>
  <DataTable :value="employees" showGridlines stripedRows paginator :rows="3" tableStyle="min-width: 50rem">
    <Column field="message_id" header="Id"></Column>
    <Column field="subject" header="Subject"></Column>
    <Column field="from" header="From"></Column>
    <Column field="to" header="To"></Column>
    <Column field="content" header="Content"></Column>
  </DataTable>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import EmailService from './service/EmailService';

onMounted(() => {
  EmailService.searchEmails("days").then((data) => (employees.value = data.emails));
});

const employees = ref([]);

</script>

<style></style>
