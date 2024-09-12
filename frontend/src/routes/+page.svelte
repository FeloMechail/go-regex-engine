<script>
  let pattern = '';
  let text = '';
  let result = null;
  let matchesarray = [];
    async function handleSubmit() {
    const response = await fetch('http://localhost:3100/api/match', {
        method: 'POST',
        headers: {
        'Content-Type': 'application/json',
        },
        body: JSON.stringify({ pattern, text }),
    });
    result = await response.json();
    matchesarray = result.matches;

}

</script>

<div class="flex flex-col items-center bg-gray-900 h-screen p-6 text-white">
    <!-- Header / Input for Regex -->
    <div class="w-full max-w-2xl">
      <label for="regex-input" class="block text-lg mb-2">REGULAR EXPRESSION</label>
      <input
        id="regex-input"
        type="text"
        bind:value={pattern}
        class="w-full p-2 bg-gray-900 border border-gray-700 shadow-md shadow-black rounded"
        placeholder="/ your regex here"
        on:input={handleSubmit}
      />
    </div>
  
    <!-- Test String -->
    <div class="w-full max-w-2xl mt-6">
        <label for="regex-text-string" class="block text-lg mb-2">TEST STRING</label>
        <textarea
            bind:value={text}
            class="w-full p-2 bg-gray-900 border border-gray-700 shadow-md shadow-black rounded h-32"
            placeholder="your test string"
            on:input={handleSubmit}
        ></textarea>
    </div>

    <!-- Result Section -->
    {#if result !== null}
    <div class="w-full max-w-2xl mt-6">
      <div class="bg-gray-700 p-4 rounded">
        {result.matched ? "Match Found: " : "No Match"}
        <ul>
          {#each matchesarray as match, index}
            <li>Match {index +2} {match.Rangee.join("-")}: {match.Match}</li>
          {/each}
        </ul>
      </div>
    </div>
  {/if}
  </div>