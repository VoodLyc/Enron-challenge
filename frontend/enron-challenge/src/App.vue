<template>
  <DataTable :value="emails" showGridlines stripedRows :size="small" lazy paginator :first="first" :rows="10"
    :totalRecords="totalRecords" :loading="loading" @page="onPage($event)" :rowsPerPageOptions="[5, 10, 20]"
    paginatorTemplate="RowsPerPageDropdown FirstPageLink PrevPageLink CurrentPageReport NextPageLink LastPageLink"
    currentPageReportTemplate="{first} to {last} of {totalRecords}">
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

const first = ref(0);
const totalRecords = ref(0);
const loading = ref(false);
const lazyParams = ref({});
const emails = ref([]);

const loadLazyData = (event) => {
  loading.value = true;
  console.log(event);
  lazyParams.value = { ...lazyParams.value, first: event?.first || first.value };
  EmailService.getEmails(lazyParams.value.first).then((data) => {
    emails.value = data.emails;
    totalRecords.value = data.total;
    loading.value = false;
  });
}

const onPage = (event) => {
  lazyParams.value = event;
  loadLazyData(event);
}


onMounted(() => {
  lazyParams.value = {
    first: 0,
    rows: 10
  };
  loadLazyData();
});

</script>

<style></style>
