FROM ghcr.io/lesomnus/dev-golang:1.22

RUN TEMP=$(mktemp -d) \
	&& cd "${TEMP}" \
	&& curl -o ./protoc.zip -sSL https://github.com/protocolbuffers/protobuf/releases/download/v27.0/protoc-27.0-linux-x86_64.zip \
	&& unzip ./protoc.zip -d /usr/local \
	&& chmod 755 /usr/local/bin/protoc \
	&& rm -rf "${TEMP}"

RUN mkdir -p /usr/local/include/google/api/ \
	&& cd /usr/local/include/google/api/ \
	&& curl -sSLO "https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/{annotations,http}.proto"



USER hypnos

RUN  echo \
		google.golang.org/protobuf/cmd/protoc-gen-go@v1.31 \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.4 \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.21.0 \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.21.0 \
	| xargs -n1 go install

# Comment below lines if you don't want to use Zsh
RUN wget https://github.com/robbyrussell/oh-my-zsh/raw/master/tools/install.sh -O - | zsh || true \
	&& git clone https://github.com/zsh-users/zsh-autosuggestions         ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-autosuggestions     \
	&& git clone https://github.com/zsh-users/zsh-syntax-highlighting.git ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-syntax-highlighting \
	&& git clone https://github.com/zsh-users/zsh-completions             ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-completions         \
	&& git clone https://github.com/supercrabtree/k                       ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/k
ADD .zshrc /home/hypnos
ENV SHELL /bin/zsh
