
<script lang="ts">
    import Pocketbase from "pocketbase"
    import { onDestroy, onMount } from "svelte";

    const pb = new Pocketbase("")
    // const uri = document.baseURI

    onMount(async () => {
        // await pb.collection("users").authWithPassword("test1@admin.com", "Notsafe@abc123")
        const res = await fetch("http://localhost:8090/api/v1/docker/c9f3a113e3df/log")
        const reader = res.body!.getReader()
        const decoder = new TextDecoder()
        while (true) {
            const {done, value} = await reader.read()
            if (done) {
                break
            }
            const text = decoder.decode(value)
            console.log(text)
        }
    })

    onDestroy(() => {
        pb.realtime.unsubscribe("docker.container.log.c9f3a113e3df")
    })

</script>

<h1>Welcome to SvelteKit</h1>
<p>Visit <a href="https://svelte.dev/docs/kit">svelte.dev/docs/kit</a> to read the documentation</p>