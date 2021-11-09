var requestOptions = {
    method: 'GET',
    redirect: 'follow'
  };
  
fetch("/api/get", requestOptions)
    .then(function(response) {
        return response.json();
        console.log(data);

      })
      .then(function(data) {
        console.log(data);
        let current_speed = data.data.current.speed;
        let current_dir = data.data.current.dir;
        let last_updated = data.data.current.created_at;

        document.getElementById("current_speed").innerHTML = current_speed + ' m/s';
        document.getElementById("current_dir").innerHTML = current_dir;
        document.getElementById("last_updated_speed").innerHTML = last_updated;
        document.getElementById("last_updated_dir").innerHTML = last_updated;

        
        let tab =
        ``;
       //  Loop to access all rows
        for (let r of data.data.log) {
            tab += `<tr>
            <th scope= "row"><a href="#">${r.created_at}</a></th>
            <td>${r.speed}m/s</td>
            <td>${r.dir}</td>		
            </tr>`;
        }
        // Setting innerHTML as tab variable
        document.getElementById("log").innerHTML = tab;


      })
    .catch(error => console.log('error', error));
