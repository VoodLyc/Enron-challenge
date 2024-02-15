<template>
  <div>
    <div class="ml-2 mb-4 mt-2 flex flex-row items-center">
      <InputGroup class="w-1/4">
        <InputText placeholder="Keyword" v-model="searchTerm" />
        <Button icon="pi pi-search" @click="onSearch()" />
      </InputGroup>
      <Button icon="pi pi-times" class="ml-3" v-if="filtered" severity="danger" rounded aria-label="Cancel" @click="unfilterData()" />
    </div>
    <DataTable :value="emails" showGridlines stripedRows lazy paginator :first="first" :rows="10"
      :totalRecords="totalRecords" :loading="loading" @page="onPage($event)" :rowsPerPageOptions="[5, 10, 20]"
      paginatorTemplate="RowsPerPageDropdown FirstPageLink PrevPageLink CurrentPageReport NextPageLink LastPageLink"
      currentPageReportTemplate="{first} to {last} of {totalRecords}">
      <Column field="message_id" header="Id"></Column>
      <Column field="subject" header="Subject"></Column>
      <Column field="from" header="From"></Column>
      <Column field="to" header="To"></Column>
      <Column field="content" header="Content"></Column>
    </DataTable>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import EmailService from './service/EmailService';

const first = ref(0);
const totalRecords = ref(0);
const loading = ref(false);
const lazyParams = ref({});
const emails = ref([]);
const searchTerm = ref("");
const filtered = ref(false);

const loadLazyData = (event) => {
  loading.value = true;
  lazyParams.value = { ...lazyParams.value, first: event?.first || first.value };
  if (filtered.value) {
    EmailService.searchEmails(searchTerm.value, lazyParams.value.first)
      .then((data) => {
        emails.value = data.emails;
        totalRecords.value = data.total;
        loading.value = false;
      })
  }
  else {
    EmailService.getEmails(lazyParams.value.first)
      .then((data) => {
        emails.value = data.emails;
        totalRecords.value = data.total;
        loading.value = false;
      });
  }
}

const onPage = (event) => {
  lazyParams.value = event;
  loadLazyData(event);
}

const onSearch = () => {
  lazyParams.value = {
    first: 0,
    rows: 10
  }
  filtered.value = true;
  loadLazyData();

}

const unfilterData = () => {
  filtered.value = false;
  lazyParams.value = {
    first: 0,
    rows: 10
  };
  loadLazyData();
}

onMounted(() => {
  lazyParams.value = {
    first: 0,
    rows: 10
  };
  loadLazyData();
});

</script>
