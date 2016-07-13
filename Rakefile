desc "install dependencies"
task "deps" do
  sh "go get ./..."
end

desc "build binary"
task "build" => [:deps] do
  sh "gox -osarch=\"linux/amd64\" -output=\"pkg/{{.OS}}_{{.Arch}}/{{.Dir}}\""
end
