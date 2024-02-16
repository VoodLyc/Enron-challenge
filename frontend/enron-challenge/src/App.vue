<template>
  <div>
    <div class="w-1/4 ml-2 mb-4 mt-2 flex flex-row items-center">
      <InputGroup>
        <InputText placeholder="Keyword" v-model="searchTerm" />
        <Button icon="pi pi-search" @click="onSearch()" />
      </InputGroup>
      <Button icon="pi pi-times" class="ml-3" v-if="filtered" severity="danger" rounded aria-label="Cancel"
        @click="onUnfilterData()" />
    </div>
    <DataTable :value="emails" showGridlines stripedRows lazy paginator :first="first.sync" :rows="5"
      :totalRecords="totalRecords" :loading="loading" @page="onPage($event)" :rowsPerPageOptions="[5, 10, 20]"
      paginatorTemplate="RowsPerPageDropdown FirstPageLink PrevPageLink CurrentPageReport NextPageLink LastPageLink"
      currentPageReportTemplate="{first} to {last} of {totalRecords}">
      <Column field="subject" header="Subject"></Column>
      <Column field="from" header="From"></Column>
      <Column field="to" header="To"></Column>
      <Column header="Content" class="w-1/2">
        <template #body="{ data }">
          <span class="whitespace-pre-line">{{ data.content }}</span>
        </template>
      </Column>
      <Column  v-if="filtered" header="Highlight">
      <template #body="{data}">
        <span v-html="data.highlight.content"></span>
      </template>
      </Column>
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
  console.log(lazyParams)
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
  filtered.value = true;
  loadLazyData();

}

const onUnfilterData = () => {
  filtered.value = false;
  loadLazyData();
}

onMounted(() => {
  lazyParams.value = {
    first: 0,
    rows: 5
  };
  loadLazyData();
});

</script>
