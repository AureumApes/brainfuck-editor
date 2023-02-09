<script lang="ts">

    function onBuild(){
        let code = document.getElementById("code-input").value
        let name = document.getElementById("name").value

        fetch("/build?name=" + name, {
            method: "post",
            body: code
        }).then()
        console.log(code)
    }
</script>

<button id="build" on:click|preventDefault={onBuild}>
    Build
</button>