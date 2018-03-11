<template>
  <div class="hero is-primary is-bold">
    <div class="hero-body">
      <div class="container">
        <h1 class="title">{{ title }}</h1>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "GcHeader",
  props: {
    title :{
      required: true,
      type: String
    }
  }
}
</script>
