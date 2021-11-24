<template>
  <div>
    <Header />
    MainPage
  </div>
</template>

<script>
import Header from '@/components/Header.vue';
import todo from '@/api/todo';

export default {
  name: 'MainPage',
  components: {
    Header,
  },
  mounted() {
    todo.fetchAll();
  },
};
</script>
