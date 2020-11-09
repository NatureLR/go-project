package template

func init() {
	Default.Add("gitlab-ci", GitlabCI, ".gitlab-ci.yml")
}

// GitlabCI 文件模板
const GitlabCI = `
# 定义全局变量
variables:
  DOCKER_REGISTRY_URL: "$REGISTRY"

# 定义阶段
stages:
  - build
  - deploy

# 所有stage脚本之前都会执行，用来定义一些要用的变量
before_script:
  - 'echo "Job $CI_JOB_NAME triggered by $GITLAB_USER_NAME ($GITLAB_USER_ID)"'
  - 'echo "before_script"'
  - 'echo "Build fold : $PWD"'
  - 'echo "PROJECT NAME : $CI_PROJECT_NAME"'
  - 'echo "PROJECT ID : $CI_PROJECT_ID"'
  - 'echo "PROJECT URL : $CI_PROJECT_URL"'
  - 'echo "DOCKER REGISTRY URL : $DOCKER_REGISTRY_URL"'

  
  - 'IMAGE_REPO="$DOCKER_REGISTRY_URL/$CI_PROJECT_NAMESPACE/$CI_PROJECT_NAME/$CI_COMMIT_REF_NAME"'
  - 'IMAGE_REPO=$(echo $IMAGE_REPO|tr "[:upper:]" "[:lower:]")'
  - 'echo IMAGE_REPO: $IMAGE_REPO'

  - 'IMAGE_TAG=$IMAGE_REPO:${CI_COMMIT_SHA:0:8}'
  - 'IMAGE_TAG_LATEST=$IMAGE_REPO:latest'

  - 'echo IMAGE_TAG is :$IMAGE_TAG'
  - 'echo IMAGE_TAG_LATEST is :$IMAGE_TAG_LATEST'
  - '====================================before_script执行完毕==========================================================='

# 定义每个阶段的按钮
build_image:
  # 绑定的stage
  stage: build
  # 只在那些分支上生效
  #only:
  #  - master
  #  - tags
  # 要用到的镜像dokcer in docker
  #image: docker:git
  # 加入的服务
  #services:
  #  - docker:18.09.7-dind
  # 执行方式为手动执行
  when: manual
  # 允许失败
  allow_failure: true
  # 一下所有脚本都是为了执行docker build
  script:
    - 'echo "Build on $CI_COMMIT_REF_NAME"'
    - 'echo "HEAD commit SHA $CI_COMMIT_SHA"'

    # 编译docker镜像 
    # TODO 改为调用makefile
    - 'docker build -f ./Dockerfile -t $IMAGE_TAG -t $IMAGE_TAG_LATEST .'
    - 'docker push $IMAGE_TAG'
    - 'docker push $IMAGE_TAG_LATEST'
    - 'echo "The build is sucessful,The image is : $IMAGE_TAG"'

deploy:
  stage: deploy 
  when: manual
  allow_failure: true
  script:
    - '<执行发布脚本>'
`
