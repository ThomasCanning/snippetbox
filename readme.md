<h1>Notes</h1>
<h2>Notes about project structure</h2>
<ul>
  <li><strong>cmd</strong> contains the application specific code for the executable applications in the project, e.g. the web application itself</li>
  <li><strong>internal</strong> contains the non-application specific code, e.g. potentially reusable code such as sql database models, code in internal can only be imported by code inside the parent of internal, i.e. within snippetbox directory</li>
  <li><strong>ui</strong> </li> contains interface assests, e.g. html templates and ui/static contains images etc</li>
  <li>Project root contains non go files, such as go.mod and make files</li>
</ul>
