<template>
  <div style="width: 100%">
  <br/>
    <h2 style="padding: 0 15px">Music list</h2>
    <div v-if="musics.length === 0" style="margin: 0 auto">
      <h3>No music to display :-(</h3>
    </div>
    <div class="row" style="margin: 0">
      <div v-for="music in musics.value" v-bind:key="music.id" class="col-sm-4" style="padding: 10px">
        <div class=" card" style="margin: 10px 0">
          <h3 class="card-header">{{ music.title }} - {{ music.artist }} |
            {{ avg(music.ratings.map(e => e.rating)).toFixed(2) }}/5
            <span class="fas fa-star" style="color: #ffff00"/></h3>
          <div class="card-body">
            <Vue3WaveAudioPlayer src="/src/assets/funk.mp3"/>
            <br/>
            <h4>Comments:</h4>
            <ul>
              <li v-for="comment in music.ratings">
                {{ comment.rating }}/5 <span class="fas fa-star"/> : {{ comment.comment }} <span
                  style="color: #aaa; font-size: 9px"> -- {{ comment.rating_date }} </span>
              </li>
              <li v-if="music.ratings.length === 0" style="color: #aaa">
                No comment to display.
              </li>
            </ul>

            <div style="display: flex">
              <input class="form-control" type="text" v-model="music.add_comment" v-on:input="$forceUpdate"/>
              <select class="form-control" v-model="music.add_rating" style="width: 60px">
                <option value="1">1</option>
                <option value="2">2</option>
                <option value="3">3</option>
                <option value="4">4</option>
                <option value="5">5</option>
              </select>
              <button type="button" class="btn btn-success" style="white-space: nowrap" v-on:click="addComment(music)"
                      :disabled="music.add_comment === ''">Add comment
              </button>
            </div>
          </div>
          <div class="card-footer text-muted">
            <div style="display: flex; justify-content: space-between">
              <div> Published on {{ music.published_date }}
                <span style="font-style: italic; padding-left: 20px; color: #aaa">(Filename : {{
                    music.file_name
                  }})</span></div>
              <div>
                <span class="fas fa-times" style="cursor: pointer; color: #ff0000" v-on:click="deleteMusic(music)"/>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <hr>
    <h2 style="padding: 0 15px">Add a music</h2>

    <div style="margin: 20px; margin-bottom: 100px">
      <form class="form" v-on:submit.prevent="addMusicToAPI()">
        <input class="form-control" type="text" placeholder="Music title" v-model="addMusic.title">
        <br />
        <input class="form-control" type="text" placeholder="Artist name" v-model="addMusic.artist">
        <br />
        <input class="form-control" type="text" placeholder="Album name" v-model="addMusic.album">
        <br />
        <input class="form-control" type="text" placeholder="Genre" v-model="addMusic.genre">
        <br />
        <input class="form-control" type="text" placeholder="File name" v-model="addMusic.file_name">
        <br />
        <button type="submit" class="btn btn-primary">Add music ></button>
      </form>
    </div>

    <!-- TODO: fix v-model -->
    <hr>
    <h2 style="padding: 0 15px">Add an album</h2>

    <div style="margin: 20px; margin-bottom: 100px">
      <form class="form" v-on:submit.prevent="addMusicToAPI()">
        <input class="form-control" type="text" placeholder="Album name" v-model="addAlbum.name">
        <br />
        <input class="form-control" type="text" placeholder="Artist name" v-model="addAlbum.artist">
        <br />
        <br />
        <button type="submit" class="btn btn-primary">Add album ></button>
      </form>
    </div>

    <hr>
    <h2 style="padding: 0 15px">Add a genre</h2>

    <div style="margin: 20px; margin-bottom: 100px">
      <form class="form" v-on:submit.prevent="addMusicToAPI()">
        <input class="form-control" type="text" placeholder="Genre" v-model="addGenre.name">
        <br />
        <button type="submit" class="btn btn-primary">Add genre ></button>
      </form>
    </div>

    <hr>
    <h2 style="padding: 0 15px">Add an artist</h2>

    <div style="margin: 20px; margin-bottom: 100px">
      <form class="form" v-on:submit.prevent="addMusicToAPI()">
        <input class="form-control" type="text" placeholder="Artist name" v-model="addArtist.name">
        <br />
        <button type="submit" class="btn btn-primary">Add artist ></button>
      </form>
    </div>
  </div>



</template>

<script setup>
import {onMounted, reactive} from "vue";
import {useAxios} from "@vueuse/integrations/useAxios";
import {useAuthStore} from "@/stores/auth_store.js";
import {storeToRefs} from "pinia";
import {useGeneralResponses} from "@/composables/general_responses.js";
import {useToast} from "vue-toast-notification";

import Vue3WaveAudioPlayer from 'vue3-wave-audio-player'


const authStore = useAuthStore()
const {currentUser} = storeToRefs(authStore)

const musics = reactive({})
const addMusic = reactive({
  title: "",
  file_name: "",
  artist: "",
  album: "",
  genre: ""
})

const artist = reactive({})
const addArtist = reactive({
  name: ""
})

const album = reactive({})
const addAlbum = reactive({
  name: "",
  artist: ""
})

const genre = reactive({})
const addGenre = reactive({
  name: ""
})

const generalResponses = useGeneralResponses()
const toast = useToast();

onMounted(() => {
  getMusics()
})

function avg(d) {
  const sum = d.reduce((a, b) => a + b, 0);
  return (sum / d.length) || 0
}

async function addMusicToAPI() {
  let dataToSend = {
    title: addMusic.title,
    file_name: addMusic.file_name,
    artist: addMusic.artist,
    album: addMusic.album,
    genre: addMusic.genre
  }

  const config = {
    headers: authStore.authAxiosConfig,
    method: 'POST',
    data: dataToSend
  }
  const {error} = await useAxios(authStore.authBaseUrl + 'musics/', config)
  if (!error.value) {
    addMusic.title = ""
    addMusic.file_name = ""
    addMusic.artist = ""
    addMusic.album = ""
    addMusic.genre = ""
    this.$forceUpdate
    toast.success("music added")
    await getMusics()
  } else {
    generalResponses.manageError(error.value)
    // manage error and let the component display it however it wants to
    return Promise.reject(error.value)
  }
}

async function deleteMusic(data) {
  const config = {
    headers: authStore.authAxiosConfig,
    method: 'DELETE',
  }
  const {error} = await useAxios(authStore.authBaseUrl + 'musics/' + data.id, config)
  if (!error.value) {
    toast.success("music deleted")
    await getMusics()
  } else {
    generalResponses.manageError(error.value)
    // manage error and let the component display it however it wants to
    return Promise.reject(error.value)
  }
}

async function addComment(data) {
  let dataToSend = {
    comment: data.add_comment,
    rating: data.add_rating
  }

  const config = {
    headers: authStore.authAxiosConfig,
    method: 'POST',
    data: dataToSend
  }
  const {error} = await useAxios(authStore.authBaseUrl + 'musics/' + data.id + "/ratings", config)
  if (!error.value) {
    data.add_comment = ""
    data.add_rating = 1
    this.$forceUpdate
    toast.success("Comment added")
    await getMusics()
  } else {
    generalResponses.manageError(error.value)
    // manage error and let the component display it however it wants to
    return Promise.reject(error.value)
  }
}

async function getMusics() {
  const config = {
    headers: authStore.authAxiosConfig,
    method: 'GET',
  }
  const {data, error} = await useAxios(authStore.authBaseUrl + 'musics/', config)
  if (!error.value) {
    data.value.forEach(e => {
      e.add_comment = "";
      e.add_rating = 1
    })
    musics.value = data
  } else {
    generalResponses.manageError(error.value)
    // manage error and let the component display it however it wants to
    return Promise.reject(error.value)
  }
}

</script>
