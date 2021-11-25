<template>
  <div>
    <Header />
    <div class="container">
      <h1 class="text-center">TODO</h1>
      <form class="row mb-5">
        <div class="col-sm-3"></div>
        <div class="col-sm-6">
          <label class="visually-hidden" for="specificSizeInputGroupUsername">Todo</label>
          <div class="input-group">
            <input
              type="text"
              class="form-control"
              id="specificSizeInputGroupUsername"
              placeholder="Todo"
            >
            <button type="button" class="btn btn-success">Создать</button>
          </div>
        </div>
        <div class="col-sm-3"></div>
      </form>
      <div class="row">
        <div class="col-3"></div>
        <div class="col-6">
          <div class="card" style="width: 100%">
            <div class="card-body">
              <h5 class="card-title">Card title</h5>
              <a href="#" class="btn btn-danger">Удалить</a>
            </div>
          </div>
        </div>
        <div class="col-3"></div>
      </div>
    </div>
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

<style scoped>
h1 {
  margin: 20px 0;
}
</style>
