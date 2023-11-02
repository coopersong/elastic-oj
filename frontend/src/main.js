window.onload = function() {
    fetch('http://localhost:8080/problems', {
        method: 'GET',
    })
        .then(response => response.json())
        .then(data => addSideBarItemsBasedOnData(data))
        .catch((error) => {
            alert("Error: " + error)
        });
}

function addSideBarItemsBasedOnData(data) {
    for (let i = 0; i < data.result.length; i++) {
        let problemID = data.result[i]["ProblemID"];
        let title = data.result[i]["Title"];

        let sideBar = document.getElementById("left");
        let listItem = document.createElement("li");
        let a = document.createElement("a");
        a.textContent = title;
        listItem.appendChild(a);
        a.setAttribute("ProblemID", problemID)
        a.addEventListener("click", function(event){
            let problemID = event.target.getAttribute("ProblemID");
            console.log(problemID)
            fetch('http://localhost:8080/problems/'+problemID, {
                method: 'GET',
            })
                .then(response => response.json())
                .then(data => refreshCodingPageBasedOnData(data))
                .catch((error) => {
                    alert("Error: " + error)
                });
        });
        sideBar.appendChild(listItem);
    }
}

function refreshCodingPageBasedOnData(data) {
    let problemID = data.result["ProblemID"]
    let title = data.result["Title"]
    let description = data.result["Description"]

    let problemDiv = document.getElementById("problem")
    problemDiv.innerHTML = ""

    let titleH = document.createElement("h1")
    titleH.textContent = title
    problemDiv.appendChild(titleH)

    let descriptionH = document.createElement("p")
    descriptionH.textContent = description
    problemDiv.appendChild(descriptionH)

    document.getElementById('submitButton').onclick = function() {
        // Retrieve the code from the textarea
        let code = document.getElementById('codeBox').value;
        fetch('http://localhost:8080/problems/run', {
            method: 'POST',
            body: JSON.stringify({
                'ProblemID': problemID,
                'SubmittedQuery': code
            }),
        })
            .then(response => response.json())
            .then(data => displayJudgeResultBasedOnData(data))
            .catch((error) => {
                console.error('Error:', error);
            });
    };
}

function displayJudgeResultBasedOnData(data) {
    if (data.message === "PASS") {
        alert("PASS")
    } else {
        alert("FAIL")
    }
}